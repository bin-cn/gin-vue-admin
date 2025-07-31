
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户昵称:" prop="nickName">
    <el-select v-model="formData.nickName" placeholder="请选择用户昵称" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.nickName" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="登录码:" prop="loginCode">
    <el-input v-model="formData.loginCode" :clearable="true" placeholder="请输入登录码" />


<el-form-item label="区服名称:" prop="gameServerName">
    <el-input v-model="formData.gameServerName" :clearable="true" placeholder="请输入区服名称" />
</el-form-item>
        <el-form-item label="区服ID:" prop="gameServerId">
    <el-input v-model.number="formData.gameServerId" :clearable="true" placeholder="请输入区服ID" />
</el-form-item>

</el-form-item>
        <el-form-item label="游戏角色名称:" prop="roleGameName">
    <el-input v-model="formData.roleGameName" :clearable="true" placeholder="请输入游戏角色名称" />
</el-form-item>
        <el-form-item label="游戏角色等级:" prop="roleGameLevel">
    <el-input v-model="formData.roleGameLevel" :clearable="true" placeholder="请输入游戏角色等级" />
</el-form-item>
        <el-form-item label="用户ID:" prop="userId">
    <el-select v-model="formData.userId" placeholder="请选择用户ID" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="未绑定元宝数量:" prop="unBoundIngotQuantity">
    <el-input v-model.number="formData.unBoundIngotQuantity" :clearable="true" placeholder="请输入未绑定元宝数量" />
</el-form-item>
        <el-form-item label="绑定元宝数量:" prop="boundIngotQuantity">
    <el-input v-model.number="formData.boundIngotQuantity" :clearable="true" placeholder="请输入绑定元宝数量" />
</el-form-item>
        <el-form-item label="元宝总数:" prop="totalIngotQuantity">
    <el-input v-model.number="formData.totalIngotQuantity" :clearable="true" placeholder="请输入元宝总数" />
</el-form-item>
        <el-form-item label="未绑定灵符:" prop="unBoundTalisman">
    <el-input v-model.number="formData.unBoundTalisman" :clearable="true" placeholder="请输入未绑定灵符" />
</el-form-item>
        <el-form-item label="绑定灵符:" prop="boundTalisman">
    <el-input v-model.number="formData.boundTalisman" :clearable="true" placeholder="请输入绑定灵符" />
</el-form-item>
        <el-form-item label="灵符总数:" prop="totalTalisman">
    <el-input v-model.number="formData.totalTalisman" :clearable="true" placeholder="请输入灵符总数" />
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
    getGameUserDataSource,
  createGameUser,
  updateGameUser,
  findGameUser
} from '@/api/smartcreate/uGameUser'

defineOptions({
    name: 'GameUserForm'
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
const formData = ref({
            nickName: '',
            loginCode: '',
            roleGameName: '',
            roleGameLevel: '',
            userId: undefined,
            unBoundIngotQuantity: undefined,
            boundIngotQuantity: undefined,
            totalIngotQuantity: undefined,
            unBoundTalisman: undefined,
            boundTalisman: undefined,
            totalTalisman: undefined,
            gameServerName: '',
            gameServerId: undefined,
        })
// 验证规则
const rule = reactive({
               nickName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
                gameServerId : [{
                required: true,
                message: '',
                trigger: ['input','blur'],
              }
            ],
               loginCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getGameUserDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGameUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
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
               res = await createGameUser(formData.value)
               break
             case 'update':
               res = await updateGameUser(formData.value)
               break
             default:
               res = await createGameUser(formData.value)
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
