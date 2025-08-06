<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="任务名称:" prop="taskName">
          <el-input v-model="formData.taskName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务分组:" prop="taskGroup">
          <el-input v-model="formData.taskGroup" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="cron_expression:" prop="cronExpression">
          <el-input v-model="formData.cronExpression" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="invokeTarget:" prop="invokeTarget">
          <el-input v-model="formData.invokeTarget" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="args:" prop="args">
          <el-input v-model="formData.args" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务执行超时时间(单位秒),0不限制:" prop="timeout">
          <el-input v-model.number="formData.timeout" :clearable="true" placeholder="please intput " />
       </el-form-item>
        <el-form-item label="是否允许多实例运行:" prop="multi">
          <el-switch v-model="formData.multi" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="重试次数:" prop="retryTimes">
          <el-switch v-model="formData.retryTimes" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="重试间隔时间:" prop="retryInterval">
          <el-input v-model.number="formData.retryInterval" :clearable="true" placeholder="please intput " />
       </el-form-item>
        <el-form-item label="tag:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="remark:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态 1:正常 0:停止:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
  createSysTask,
  updateSysTask,
  findSysTask
} from '@/plugin/sysTask/api/sysTask'

defineOptions({
    name: 'SysTaskForm'
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
            taskName: '',
            taskGroup: '',
            cronExpression: '',
            invokeTarget: '',
            args: '',
            timeout: 0,
            multi: 0,
            retryTimes: 0,
            retryInterval: 0,
            tag: '',
            remark: '',
            status: 1,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resysTask
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
               res = await createSysTask(formData.value)
               break
             case 'update':
               res = await updateSysTask(formData.value)
               break
             default:
               res = await createSysTask(formData.value)
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
