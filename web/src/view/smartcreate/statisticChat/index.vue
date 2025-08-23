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
import { ref, onMounted, nextTick } from 'vue'
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

// 响应式数据
const chartData = ref<DailyRevenueStatistic[]>([])
const currentDate = ref('')
const activeTab = ref('user')
const userChartRef = ref()
const serverChartRef = ref()

function handleTabChange (tabName: string) {
  if (tabName === 'user') {
    loadChartData(currentDate.value, 1).then(() => {
      // 延迟重新渲染，确保DOM完全更新
      setTimeout(() => {
        nextTick(() => {
          if (userChartRef.value) {
            userChartRef.value.resizeChart()
          }
        })
      }, 100)
    })
  } else if (tabName === 'server') {
    loadChartData(currentDate.value, 2).then(() => {
      // 延迟重新渲染，确保DOM完全更新
      setTimeout(() => {
        nextTick(() => {
          if (serverChartRef.value) {
            serverChartRef.value.resizeChart()
          }
        })
      }, 100)
    })
  }
}

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

const getCurrentYearMonth = () => {
  const now = new Date()
  const y = now.getFullYear()
  const m = String(now.getMonth() + 1).padStart(2, '0')
  return `${y}-${m}`
}

onMounted(() => {
  currentDate.value = getCurrentYearMonth()
  loadChartData(currentDate.value, 1)
})
</script>