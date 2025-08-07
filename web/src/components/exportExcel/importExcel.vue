<template>
  <el-upload
    v-if="!customImport"
    :action="url"
    :show-file-list="false"
    :on-success="handleSuccess"
    :multiple="false"
    :headers="{'x-token': token}"
  >
    <el-button type="primary" icon="upload" class="ml-3"> 导入 </el-button>
  </el-upload>
  <el-upload
    v-else
    :action="''"
    :show-file-list="false"
    :http-request="customUpload"
    :before-upload="beforeUpload"
    :multiple="false"
    :headers="{'x-token': token}"
  >
    <el-button type="primary" icon="upload" class="ml-3"> 导入 </el-button>
  </el-upload>
</template>

<script setup>
  import { ElMessage } from 'element-plus'
  import { useUserStore } from "@/pinia";
  import { importGameUserAuthCodeCustom } from '@/api/smartcreate/uGameUserAuthCodeImport'

  let baseUrl = import.meta.env.VITE_BASE_API
  if (baseUrl === "/"){
    baseUrl = ""
  }

  const props = defineProps({
    templateId: {
      type: String,
      required: true
    },
    customImport: {
      type: Boolean,
      default: false
    }
  })

  const userStore = useUserStore()

  const token = userStore.token

  const emit = defineEmits(['on-success'])

  const url = `${baseUrl}/sysExportTemplate/importExcel?templateID=${props.templateId}`

  const handleSuccess = (res) => {
    if (res.code === 0) {
      ElMessage.success('导入成功')
      emit('on-success')
    } else {
      ElMessage.error(res.msg)
    }
  }

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

  const customUpload = (options) => {
    const { file } = options
    const formData = new FormData()
    formData.append('file', file)
    
    importGameUserAuthCodeCustom(file)
      .then(response => {
        if (response.code === 0) {
          ElMessage.success('导入成功')
          emit('on-success')
        } else {
          ElMessage.error(response.msg)
        }
      })
      .catch(error => {
        ElMessage.error('导入失败')
        console.error('Import error:', error)
      })
  }
</script>
