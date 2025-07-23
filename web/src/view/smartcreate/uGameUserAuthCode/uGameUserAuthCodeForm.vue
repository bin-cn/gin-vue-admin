
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="登录码:" prop="loginCode">
    <el-input v-model="formData.loginCode" :clearable="true" placeholder="请输入登录码" />
</el-form-item>
        <el-form-item label="机器编号:" prop="machineNoName">
    <el-input v-model="formData.machineNoName" :clearable="true" placeholder="请输入机器编号" />
</el-form-item>
        <el-form-item label="使用人:" prop="assignerName">
    <el-input v-model="formData.assignerName" :clearable="true" placeholder="请输入使用人" />
</el-form-item>
        <el-form-item label="账号:" prop="account">
    <el-input v-model="formData.account" :clearable="true" placeholder="请输入账号" />
</el-form-item>
        <el-form-item label="密码:" prop="password">
    <el-input v-model="formData.password" :clearable="true" placeholder="请输入密码" />
</el-form-item>
        <el-form-item label="区服:" prop="gameServerName">
    <el-input v-model="formData.gameServerName" :clearable="true" placeholder="请输入区服" />
</el-form-item>
        <el-form-item label="区服id:" prop="gameServerId">
    <el-input v-model.number="formData.gameServerId" :clearable="true" placeholder="请输入区服id" />
</el-form-item>
        <el-form-item label="游戏角色名字:" prop="roleGameName">
    <el-input v-model="formData.roleGameName" :clearable="true" placeholder="请输入游戏角色名字" />
</el-form-item>
        <el-form-item label="游戏角色ID:" prop="roleGameId">
    <el-input v-model.number="formData.roleGameId" :clearable="true" placeholder="请输入游戏角色ID" />
</el-form-item>
        <el-form-item label="开区时间:" prop="serverOpenTime">
    <el-input v-model="formData.serverOpenTime" :clearable="true" placeholder="请输入开区时间" />
</el-form-item>
        <el-form-item label="进区时间:" prop="enterServerTime">
    <el-input v-model="formData.enterServerTime" :clearable="true" placeholder="请输入进区时间" />
</el-form-item>
        <el-form-item label="账号状态:" prop="accountStatus">
    <el-select v-model="formData.accountStatus" placeholder="请选择账号状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in AccountStatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
</el-form-item>
        <el-form-item label="ID名字:" prop="iDName">
    <el-input v-model="formData.iDName" :clearable="true" placeholder="请输入ID名字" />
</el-form-item>
        <el-form-item label="身份证号码:" prop="iDCardNumber">
    <el-input v-model="formData.iDCardNumber" :clearable="true" placeholder="请输入身份证号码" />
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
  createGameUserAuthCode,
  updateGameUserAuthCode,
  findGameUserAuthCode
} from '@/api/smartcreate/uGameUserAuthCode'

defineOptions({
    name: 'GameUserAuthCodeForm'
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
const AccountStatusOptions = ref([])
const formData = ref({
            loginCode: '',
            machineNoName: '',
            assignerName: '',
            account: '',
            password: '',
            gameServerName: '',
            gameServerId: undefined,
            roleGameName: '',
            roleGameId: undefined,
            serverOpenTime: '',
            enterServerTime: '',
            accountStatus: '',
            remark: '',
            iDName: '',
            iDCardNumber: '',
        })
// 验证规则
const rule = reactive({
               loginCode : [{
                   required: true,
                   message: '登录码错误或重复,请检查再导入',
                   trigger: ['input','blur'],
               }],
               machineNoName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               assignerName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               account : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gameServerName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gameServerId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roleGameName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               roleGameId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               iDName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               iDCardNumber : [{
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
      const res = await findGameUserAuthCode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    AccountStatusOptions.value = await getDictFunc('AccountStatus')
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
               res = await createGameUserAuthCode(formData.value)
               break
             case 'update':
               res = await updateGameUserAuthCode(formData.value)
               break
             default:
               res = await createGameUserAuthCode(formData.value)
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
