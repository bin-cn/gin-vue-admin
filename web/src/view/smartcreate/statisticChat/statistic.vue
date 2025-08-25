<template>
  <div>
    <el-tabs v-model="activeTab" @tab-change="handleTabChange">
      <el-tab-pane label="用户统计" name="user">
        <UserStatChart ref="userChartRef" :data="chartData" />
      </el-tab-pane>
      <el-tab-pane label="区服统计" name="server">
        <ServerZoneChart ref="serverChartRef" :data="chartData" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import UserStatChart from './components/UserStatChart.vue'
import ServerZoneChart from './components/ServerZoneChart.vue'
import { getStatistic } from '@/api/smartcreate/dailyrevenuerecord'

// 定义数据类型（与后端一致）
interface DailyRevenueStatistic {
  statisticDate: string
  userId: string
  userNickname: string
  serverName: string
  serverZoneId: string
  statisticType: string
  amount: number
}

// 路由相关
const route = useRoute()
const router = useRouter()

// 响应式数据
const chartData = ref<DailyRevenueStatistic[]>([])
const currentDate = ref('')
const activeTab = ref('user')
const userChartRef = ref()
const serverChartRef = ref()

// 加载图表数据的函数
const loadChartData = async (date: string, type: number) => {
  try {
    const res = await getStatistic({
      statisticDate: date,
      statisticType: type
    })
    chartData.value = res.data || []
  } catch (e) {
    console.error('获取统计数据失败', e)
  }
}

// 获取当前年月
const getCurrentYearMonth = () => {
  const now = new Date()
  const y = now.getFullYear()
  const m = String(now.getMonth() + 1).padStart(2, '0')
  return `${y}-${m}`
}

// 处理标签页变化
const handleTabChange = (tabName: string) => {
  // 更新URL参数
  router.replace({
    query: { ...route.query, tab: tabName }
  })

  const type = tabName === 'user' ? 1 : 2
  loadChartData(currentDate.value, type).then(() => {
    // 使用nextTick确保DOM更新完成后再调用resize
    nextTick(() => {
      setTimeout(() => {
        if (tabName === 'user' && userChartRef.value) {
          userChartRef.value.resizeChart?.()
        } else if (tabName === 'server' && serverChartRef.value) {
          serverChartRef.value.resizeChart?.()
        }
      }, 100)
    })
  })
}

// 监听路由参数变化
watch(() => route.query.tab, (newTab) => {
  if (newTab === 'user' || newTab === 'server') {
    activeTab.value = newTab
  }
}, { immediate: false })

// 初始化
onMounted(() => {
  currentDate.value = getCurrentYearMonth()
  
  // 初始化时根据路由参数设置标签页
  const tabParam = route.query.tab as string
  if (tabParam === 'user' || tabParam === 'server') {
    activeTab.value = tabParam
  }
  
  const type = activeTab.value === 'user' ? 1 : 2
  loadChartData(currentDate.value, type)
})
</script>