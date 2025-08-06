<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="任务id:" prop="taskId">
          <el-input v-model.number="formData.taskId" :clearable="true" placeholder="please intput " />
       </el-form-item>
        <el-form-item label="任务名称:" prop="taskName">
          <el-input v-model="formData.taskName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="cron_expression:" prop="cronExpression">
          <el-input v-model="formData.cronExpression" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="invoke_target:" prop="invokeTarget">
          <el-input v-model="formData.invokeTarget" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="timeout:" prop="timeout">
          <el-input v-model.number="formData.timeout" :clearable="true" placeholder="please intput " />
       </el-form-item>
        <el-form-item label="retryTimes:" prop="retryTimes">
          <el-switch v-model="formData.retryTimes" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="startTime:" prop="startTime">
          <el-date-picker v-model="formData.startTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="endTime:" prop="endTime">
          <el-date-picker v-model="formData.endTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="status:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="result:" prop="result">
          <el-input v-model="formData.result" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="totalTime:" prop="totalTime">
          <el-input v-model.number="formData.totalTime" :clearable="true" placeholder="please intput " />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">Save</el-button>
          <el-button type="primary" @click="back">Back</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSysTaskLog,
  updateSysTaskLog,
  findSysTaskLog
} from '@/plugin/sysTask/api/sysTaskLog'

defineOptions({
    name: 'SysTaskLogForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            taskId: 0,
            taskName: '',
            cronExpression: '',
            invokeTarget: '',
            timeout: 0,
            retryTimes: false,
            startTime: new Date(),
            endTime: new Date(),
            status: false,
            result: '',
            totalTime: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysTaskLog({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resysTaskLog
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSysTaskLog(formData.value)
               break
             case 'update':
               res = await updateSysTaskLog(formData.value)
               break
             default:
               res = await createSysTaskLog(formData.value)
               break
           }
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
