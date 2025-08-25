<template>
  <div>
    <div style="position: relative;">
      <div ref="chartRef" style="width: 100%; height: 600px; min-width: 800px;"></div>
      <el-button 
        type="primary" 
        circle 
        style="position: absolute; top: 10px; right: 10px; z-index: 100;"
        @click="toggleFullscreen"
        :title="isFullscreen ? '退出全屏' : '全屏显示'">
        <el-icon>
          <FullScreen v-if="!isFullscreen" />
          <Close v-else />
        </el-icon>
      </el-button>
    </div>
    <div v-show="!isFullscreen" style="margin-top: 20px; padding: 10px; border: 1px solid #eee; border-radius: 4px;">
      <h4 style="margin-bottom: 10px;">选择要显示的用户：</h4>
      <el-checkbox-group v-model="selectedUserKeys" @change="updateChart">
        <el-checkbox 
          v-for="user in userList" 
          :key="user.key" 
          :value="user.key"
          style="margin-right: 15px; margin-bottom: 5px;">
          {{ user.nickname }} (ID: {{ user.userId }})
        </el-checkbox>
      </el-checkbox-group>
    </div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import { onMounted, ref, watch, computed, nextTick } from 'vue'
import { FullScreen, Close } from '@element-plus/icons-vue'

interface DailyRevenueStatistic {
  statisticDate: string
  userId: string
  userNickname: string
  serverName: string
  serverZoneId: string
  statisticType: string
  amount: number
}

const props = defineProps<{
  data: DailyRevenueStatistic[]
}>()

const chartRef = ref()
let chartInstance: echarts.ECharts | null = null
const selectedUserKeys = ref<string[]>([])
const isFullscreen = ref(false)

// 全屏功能
const toggleFullscreen = async () => {
  if (!chartRef.value) return

  const chartContainer = chartRef.value.parentElement
  
  if (!isFullscreen.value) {
    // 进入全屏
    try {
      if (chartContainer.requestFullscreen) {
        await chartContainer.requestFullscreen()
      } else if ((chartContainer as any).webkitRequestFullscreen) {
        await (chartContainer as any).webkitRequestFullscreen()
      } else if ((chartContainer as any).mozRequestFullScreen) {
        await (chartContainer as any).mozRequestFullScreen()
      } else if ((chartContainer as any).msRequestFullscreen) {
        await (chartContainer as any).msRequestFullscreen()
      }

       const chartDivFullscreen = chartRef.value
       chartDivFullscreen.style.width = '100%'
       chartDivFullscreen.style.height = '100%'
       chartDivFullscreen.style.minWidth = '100%'
    } catch (error) {
      console.error('进入全屏失败:', error)
      // 如果浏览器不支持全屏API，使用CSS全屏模式
      enterCSSFullscreen()
    }
  } else {
    // 退出全屏
    try {
      if (document.exitFullscreen) {
        await document.exitFullscreen()
      } else if ((document as any).webkitExitFullscreen) {
        await (document as any).webkitExitFullscreen()
      } else if ((document as any).mozCancelFullScreen) {
        await (document as any).mozCancelFullScreen()
      } else if ((document as any).msExitFullscreen) {
        await (document as any).msExitFullscreen()
      }

       const chartDivFullscreen = chartRef.value
       chartDivFullscreen.style.width = '100%'
       chartDivFullscreen.style.height = '600px'
       chartDivFullscreen.style.minWidth = '100%'
    } catch (error) {
      console.error('退出全屏失败:', error)
      exitCSSFullscreen()
    }
  }
}

// CSS全屏模式
const enterCSSFullscreen = () => {
  const chartContainer = chartRef.value.parentElement
  chartContainer.style.position = 'fixed'
  chartContainer.style.top = '0'
  chartContainer.style.left = '0'
  chartContainer.style.width = '100vw'
  chartContainer.style.height = '100vh'
  chartContainer.style.zIndex = '9999'
  chartContainer.style.backgroundColor = 'white'
  
  // 设置图表容器占满全屏

  
  // 调整图表大小
  nextTick(() => {
    chartInstance?.resize()
  })
  
  isFullscreen.value = true
}

const exitCSSFullscreen = () => {
  const chartContainer = chartRef.value.parentElement
  chartContainer.style.position = ''
  chartContainer.style.top = ''
  chartContainer.style.left = ''
  chartContainer.style.width = ''
  chartContainer.style.height = ''
  chartContainer.style.zIndex = ''
  chartContainer.style.backgroundColor = ''
  
  // 恢复图表容器原始样式
  const chartDiv = chartRef.value
  chartDiv.style.width = '100%'
  chartDiv.style.height = '600px'
  chartDiv.style.minWidth = '800px'
  
  // 恢复原始大小
  nextTick(() => {
    chartInstance?.resize()
  })
  
  isFullscreen.value = false
}

// 监听全屏状态变化
const handleFullscreenChange = () => {
  isFullscreen.value = !!(
    document.fullscreenElement ||
    (document as any).webkitFullscreenElement ||
    (document as any).mozFullScreenElement ||
    (document as any).msFullscreenElement
  )
  
  // 全屏状态变化时重新调整图表大小
  nextTick(() => {
    chartInstance?.resize()
  })
}

// 监听ESC键退出全屏
const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && isFullscreen.value) {
    toggleFullscreen()
  }
}

onMounted(() => {
  document.addEventListener('fullscreenchange', handleFullscreenChange)
  document.addEventListener('webkitfullscreenchange', handleFullscreenChange)
  document.addEventListener('mozfullscreenchange', handleFullscreenChange)
  document.addEventListener('MSFullscreenChange', handleFullscreenChange)
  document.addEventListener('keydown', handleKeyDown)
  
  // 监听窗口大小变化
  window.addEventListener('resize', handleResize)
})

// 窗口大小变化处理
const handleResize = () => {
  nextTick(() => {
    chartInstance?.resize()
  })
}

// 清理事件监听
watch(() => chartInstance, (newInstance) => {
  if (!newInstance) {
    document.removeEventListener('fullscreenchange', handleFullscreenChange)
    document.removeEventListener('webkitfullscreenchange', handleFullscreenChange)
    document.removeEventListener('mozfullscreenchange', handleFullscreenChange)
    document.removeEventListener('MSFullscreenChange', handleFullscreenChange)
    document.removeEventListener('keydown', handleKeyDown)
    window.removeEventListener('resize', handleResize)
  }
})

const formatAmount = (amount: number): string => {
  const value = Math.abs(amount)
  
  if (value >= 100000000) {
    const yi = Math.floor(value / 100000000)
    const wan = Math.floor((value % 100000000) / 10000)
    if (wan > 0) {
      return `${yi}亿${wan}万`
    } else {
      return `${yi}亿`
    }
  } else if (value >= 10000) {
    const wan = Math.floor(value / 10000)
    return `${wan}万`
  } else {
    return `${value}`
  }
}

const userList = computed(() => {
  const userTotals = new Map<string, { userId: string; nickname: string; totalAmount: number }>()
  
  props.data.forEach(item => {
    const key = `${item.userId}_${item.userNickname}`
    const current = userTotals.get(key) || { userId: item.userId, nickname: item.userNickname, totalAmount: 0 }
    current.totalAmount += item.amount
    userTotals.set(key, current)
  })

  return Array.from(userTotals.entries())
    .map(([key, data]) => ({ key, ...data }))
    .sort((a, b) => b.totalAmount - a.totalAmount)
})

// 默认选择金额最大的用户
watch(() => userList.value, (newList) => {
  if (newList.length > 0 && selectedUserKeys.value.length === 0) {
    selectedUserKeys.value = [newList[0].key]
  }
}, { immediate: true })

const renderChart = () => {
  if (!chartRef.value || !props.data.length) return

  // 检查DOM尺寸
  const container = chartRef.value
  if (!container || container.clientWidth === 0 || container.clientHeight === 0) {
    //console.warn('ECharts容器尺寸为0，延迟重试...')

    // setTimeout(() => {
    //   renderChart()
    // }, 100)


    return
  }

  // 如果已有实例，先销毁
  if (chartInstance) {
    chartInstance.dispose()
    chartInstance = null
  }

  // 初始化图表
  chartInstance = echarts.init(chartRef.value)

  // 获取所有唯一的日期，按日期排序
  const allDates = Array.from(new Set(props.data.map(item => item.statisticDate)))
    .sort((a, b) => {
      const numA = parseInt(a)
      const numB = parseInt(b)
      return numA - numB
    })

  // 按用户分组数据，使用userId_userNickname作为key
  const userData = new Map<string, Map<string, number>>()
  
  props.data.forEach(item => {
    const userKey = `${item.userId}_${item.userNickname}`
    if (!userData.has(userKey)) {
      userData.set(userKey, new Map())
    }
    userData.get(userKey)!.set(item.statisticDate, item.amount)
  })

  // 只显示选中的用户
  const selectedUserData = Array.from(userData.entries())
    .filter(([userKey]) => selectedUserKeys.value.includes(userKey))

  if (selectedUserData.length === 0) {
    chartInstance.setOption({
      title: {
        text: '暂无数据',
        left: 'center',
        top: 'center'
      }
    }, { notMerge: true })
    return
  }

  const series = selectedUserData.map(([userKey, dateData]) => {
    const [userId, userNickname] = userKey.split('_')
    return {
      name: `${userNickname}(${userId})`,
      type: 'line',
      smooth: true,
      symbol: 'circle',
      symbolSize: 6,
      lineStyle: {
        width: 2
      },
      emphasis: {
        focus: 'series',
        lineStyle: {
          width: 4
        }
      },
      data: allDates.map(date => ({
        value: dateData.get(date) || 0,
        name: date
      }))
    }
  })

  const option = {
    title: {
      text: '用户收入统计',
      left: 'center',
      textStyle: {
        fontSize: 16,
        fontWeight: 'bold'
      }
    },
    tooltip: {
      trigger: 'axis',
      backgroundColor: 'rgba(50,50,50,0.9)',
      borderColor: '#333',
      textStyle: {
        color: '#fff'
      },
      formatter: function(params: any) {
        let result = params[0].name + '<br/>'
        params.forEach((item: any) => {
          result += item.marker + item.seriesName + ': ' + formatAmount(item.value) + '<br/>'
        })
        return result
      }
    },
    legend: {
      data: selectedUserData.map(([userKey]) => {
        const [userId, userNickname] = userKey.split('_')
        return `${userNickname}(${userId})`
      }),
      top: 40,
      type: 'scroll',
      orient: 'horizontal',
      textStyle: {
        fontSize: 12
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '15%',
      top: '15%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: allDates,
      axisLabel: {
        rotate: 45,
        fontSize: 11
      }
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        formatter: (value: number) => formatAmount(value),
        fontSize: 11
      },
      splitLine: {
        lineStyle: {
          type: 'dashed'
        }
      }
    },
    dataZoom: [
      {
        type: 'slider',
        show: true,
        xAxisIndex: [0],
        start: 0,
        end: 100,
        height: 25,
        bottom: 5,
        handleSize: '110%',
        handleStyle: {
          color: '#409EFF'
        }
      },
      {
        type: 'inside',
        xAxisIndex: [0],
        start: 0,
        end: 100
      }
    ],
    series
  }

  chartInstance.setOption(option, { notMerge: true })
  
  // 确保图表正确适应容器大小
  nextTick(() => {
    chartInstance?.resize()
  })
}

const updateChart = () => {
  renderChart()
}

onMounted(() => {
  renderChart()
})

watch(() => props.data, () => {
  renderChart()
})
</script>