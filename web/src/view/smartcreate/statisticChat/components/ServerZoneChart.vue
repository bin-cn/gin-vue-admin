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
      <h4 style="margin-bottom: 10px;">选择要显示的区服：</h4>
      
      <!-- 大区复选框 -->
      <div style="margin-bottom: 15px; padding-bottom: 15px; border-bottom: 1px solid #eee;">
        <h5 style="margin-bottom: 10px; color: #409EFF;">按大区选择：</h5>
        <el-checkbox-group v-model="selectedRegions" @change="handleRegionChange">
          <el-checkbox 
            v-for="region in regionList" 
            :key="region.name" 
            :value="region.name"
            style="margin-right: 15px; margin-bottom: 5px;">
            {{ region.name }} (总额: {{ formatAmount(region.totalAmount) }})
          </el-checkbox>
        </el-checkbox-group>
      </div>

      <!-- 区服复选框 -->
      <div>
        <h5 style="margin-bottom: 10px; color: #67C23A;">具体区服：</h5>
        <el-checkbox-group v-model="selectedServers" @change="updateChart">
          <el-checkbox 
            v-for="server in serverList" 
            :key="server.name" 
            :value="server.name"
            style="margin-right: 15px; margin-bottom: 5px;">
            {{ server.name }} (总额: {{ formatAmount(server.totalAmount) }})
          </el-checkbox>
        </el-checkbox-group>
      </div>
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
const selectedServers = ref<string[]>([])
const selectedRegions = ref<string[]>([])
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
  const chartDiv = chartRef.value
  chartDiv.style.width = '100%'
  chartDiv.style.height = '100%'
  chartDiv.style.minWidth = '100%'
  
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

// 格式化金额显示为xxx亿xxx万
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

// 计算每个区服的总金额（按名称排序）
const serverList = computed(() => {
  const serverTotals = new Map<string, number>()
  
  props.data.forEach(item => {
    const current = serverTotals.get(item.serverName) || 0
    serverTotals.set(item.serverName, current + item.amount)
  })

  return Array.from(serverTotals.entries())
    .map(([name, totalAmount]) => ({ name, totalAmount }))
    .sort((a, b) => {
      // 先按大区名排序，再按完整名称排序
      const regionA = a.name.substring(0, 2)
      const regionB = b.name.substring(0, 2)
      
      if (regionA !== regionB) {
        return regionA.localeCompare(regionB, 'zh-CN')
      }
      
      // 同一大区内按数字排序（如果包含数字）
      const numA = parseInt(a.name.substring(2)) || 0
      const numB = parseInt(b.name.substring(2)) || 0
      
      if (numA !== numB) {
        return numA - numB
      }
      
      // 最后按完整字符串排序
      return a.name.localeCompare(b.name, 'zh-CN')
    })
})

// 计算每个大区的总金额和包含的区服
const regionList = computed(() => {
  const regionMap = new Map<string, { totalAmount: number; servers: string[] }>()
  
  props.data.forEach(item => {
    const regionName = item.serverName.substring(0, 2) // 取前两个字作为大区名
    const current = regionMap.get(regionName) || { totalAmount: 0, servers: [] }
    
    // 累加该区服金额
    const serverCurrent = current.totalAmount + item.amount
    
    // 确保区服名不重复
    if (!current.servers.includes(item.serverName)) {
      current.servers.push(item.serverName)
    }
    
    regionMap.set(regionName, {
      totalAmount: serverCurrent,
      servers: current.servers
    })
  })

  return Array.from(regionMap.entries())
    .map(([name, data]) => ({ name, totalAmount: data.totalAmount, servers: data.servers }))
    .sort((a, b) => b.totalAmount - a.totalAmount)
})

// 处理大区选择变化
const handleRegionChange = (selectedRegions: string[]) => {
  // 获取所有选中大区下的区服
  const regionServers: string[] = []
  selectedRegions.forEach(regionName => {
    const region = regionList.value.find(r => r.name === regionName)
    if (region) {
      regionServers.push(...region.servers)
    }
  })
  
  // 更新选中的区服（去重）
  selectedServers.value = Array.from(new Set(regionServers))
  updateChart()
}

// 默认选择金额最大的区服
watch(() => serverList.value, (newList) => {
  if (newList.length > 0 && selectedServers.value.length === 0) {
    selectedServers.value = [newList[0].name]
    
    // 同时选中对应的大区
    const topServer = newList[0].name
    const regionName = topServer.substring(0, 2)
    const regionExists = regionList.value.find(r => r.name === regionName)
    if (regionExists && regionExists.servers.includes(topServer)) {
      selectedRegions.value = [regionName]
    }
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

  // 按区服分组数据
  const serverData = new Map<string, Map<string, number>>()
  
  props.data.forEach(item => {
    if (!serverData.has(item.serverName)) {
      serverData.set(item.serverName, new Map())
    }
    serverData.get(item.serverName)!.set(item.statisticDate, item.amount)
  })

  // 只显示选中的区服
  const selectedServerData = Array.from(serverData.entries())
    .filter(([serverName]) => selectedServers.value.includes(serverName))

  if (selectedServerData.length === 0) {
    chartInstance.setOption({
      title: {
        text: '暂无数据',
        left: 'center',
        top: 'center'
      }
    }, { notMerge: true })
    return
  }

  const series = selectedServerData.map(([serverName, dateData]) => ({
    name: serverName,
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
  }))

  const option = {
    title: {
      text: '区服收入统计',
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
      data: selectedServerData.map(([serverName]) => serverName),
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