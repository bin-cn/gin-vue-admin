
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="统计日期:" prop="statisticDate">
    <el-input v-model="formData.statisticDate" :clearable="false" placeholder="请输入统计日期" />
</el-form-item>
        <el-form-item label="用户ID:" prop="userId">
    <el-input v-model="formData.userId" :clearable="false" placeholder="请输入用户ID" />
</el-form-item>
        <el-form-item label="用户昵称:" prop="userNickname">
    <el-input v-model="formData.userNickname" :clearable="false" placeholder="请输入用户昵称" />
</el-form-item>
        <el-form-item label="区服名字:" prop="serverName">
    <el-input v-model="formData.serverName" :clearable="false" placeholder="请输入区服名字" />
</el-form-item>
        <el-form-item label="区服ID:" prop="serverZoneId">
    <el-input v-model="formData.serverZoneId" :clearable="false" placeholder="请输入区服ID" />
</el-form-item>
        <el-form-item label="主区服ID:" prop="mainServerZoneId">
    <el-input v-model="formData.mainServerZoneId" :clearable="false" placeholder="请输入主区服ID" />
</el-form-item>
        <el-form-item label="统计类型:" prop="statisticType">
    <el-select v-model="formData.statisticType" placeholder="请选择统计类型" style="width:100%" filterable :clearable="false">
        <el-option v-for="(item,key) in MaterialTypeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="产出金额:" prop="amount">
    <el-input-number v-model="formData.amount" style="width:100%" :precision="2" :clearable="false" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createDailyRevenueRecord,
  updateDailyRevenueRecord,
  findDailyRevenueRecord
} from '@/api/smartcreate/dailyrevenuerecord'

defineOptions({
    name: 'DailyRevenueRecordForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const MaterialTypeOptions = ref([])
const formData = ref({
            statisticDate: '',
            userId: '',
            userNickname: '',
            serverName: '',
            serverZoneId: '',
            mainServerZoneId: '',
            statisticType: '',
            amount: 0,
        })
// 验证规则
const rule = reactive({
               statisticDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userNickname : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverZoneId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               mainServerZoneId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               statisticType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               amount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDailyRevenueRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    MaterialTypeOptions.value = await getDictFunc('MaterialType')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createDailyRevenueRecord(formData.value)
               break
             case 'update':
               res = await updateDailyRevenueRecord(formData.value)
               break
             default:
               res = await createDailyRevenueRecord(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
