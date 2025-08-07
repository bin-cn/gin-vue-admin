<template>
  <div>
    <el-upload
      :action="url"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleError"
      :before-upload="beforeUpload"
      :multiple="false"
      :headers="{'x-token': token}"
    >
      <el-button type="primary" icon="upload" class="ml-3"> 导入 </el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { ElMessage } from 'element-plus'
import { useUserStore } from "@/pinia"
import { importGameUserAuthCodeCustom } from '@/api/smartcreate/uGameUserAuthCodeImport'

const props = defineProps({
  onSuccess: {
    type: Function,
    default: () => {}
  }
})

const userStore = useUserStore()
const token = userStore.token

const beforeUpload = (file) => {
  const isExcel = file.type === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' || 
                  file.type === 'application/vnd.ms-excel'
  if (!isExcel) {
    ElMessage.error('只能上传Excel文件!')
    return false
  }
  const isLt10M = file.size / 1024 / 1024 < 10
  if (!isLt10M) {
    ElMessage.error('文件大小不能超过 10MB!')
    return false
  }
  return true
}

const url = '' // 空字符串，因为我们使用自定义上传

const handleSuccess = (response, file) => {
  // 这里不需要，因为我们使用自定义上传方法
}

const handleError = (error) => {
  ElMessage.error('上传失败')
  console.error('Upload error:', error)
}

const customUpload = (file) => {
  importGameUserAuthCodeCustom(file.raw || file)
    .then(response => {
      if (response.code === 0) {
        ElMessage.success('导入成功')
        props.onSuccess()
      } else {
        ElMessage.error(response.msg)
      }
    })
    .catch(error => {
      ElMessage.error('导入失败')
      console.error('Import error:', error)
    })
  return false // 阻止自动上传
}
</script>