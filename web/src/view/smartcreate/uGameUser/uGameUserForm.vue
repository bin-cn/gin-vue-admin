
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
       <el-form-item label="角色在线状态:" prop="roleOnlineStatus">
    <el-select v-model="formData.roleOnlineStatus" placeholder="请选择角色在线状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in OnlineStatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="脚本在线状态:" prop="scriptOnlineStatus">
    <el-select v-model="formData.scriptOnlineStatus" placeholder="请选择脚本在线状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in OnlineStatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="封号状态:" prop="bannedStatus">
    <el-select v-model="formData.bannedStatus" placeholder="请选择封号状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in BannedStatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="灵符变动差异:" prop="talismanDiff">
    <el-input v-model.number="formData.talismanDiff" :clearable="true" placeholder="请输入灵符变动差异" />
</el-form-item>
        <el-form-item label="元宝变动差异:" prop="ingotDiff">
    <el-input v-model.number="formData.ingotDiff" :clearable="true" placeholder="请输入元宝变动差异" />
</el-form-item>
        <el-form-item label="最后一次同步查询时间:" prop="lastSyncQueryTime">
    <el-date-picker v-model="formData.lastSyncQueryTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="最后一次同步更新时间:" prop="lastSyncUpdateTime">
    <el-date-picker v-model="formData.lastSyncUpdateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="脚本最后在线时间:" prop="scriptLastOnlineTime">
    <el-date-picker v-model="formData.scriptLastOnlineTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="角色最后在线时间:" prop="roleLastOnlineTime">
    <el-date-picker v-model="formData.roleLastOnlineTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="游戏角色ID:" prop="roleGameId">
    <el-input v-model="formData.roleGameId" :clearable="true" placeholder="请输入游戏角色ID" />
</el-form-item>
        <el-form-item label="原始区服ID:" prop="serverZoneId">
    <el-input v-model="formData.serverZoneId" :clearable="true" placeholder="请输入原始区服ID" />
</el-form-item>
        <el-form-item label="最后一次交易元宝变动时间:" prop="lastIngotTradeTime">
    <el-date-picker v-model="formData.lastIngotTradeTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="最后一次交易灵符时间:" prop="lastTalismanTradeTime">
    <el-date-picker v-model="formData.lastTalismanTradeTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="线上灵符总数:" prop="onlineTalismanTotal">
    <el-input v-model.number="formData.onlineTalismanTotal" :clearable="true" placeholder="请输入线上灵符总数" />
</el-form-item>
        <el-form-item label="线上元宝总数:" prop="onlineIngotTotal">
    <el-input v-model.number="formData.onlineIngotTotal" :clearable="true" placeholder="请输入线上元宝总数" />
</el-form-item>
        <el-form-item label="最后一次线上查询时间:" prop="lastOnlineQueryTime">
    <el-date-picker v-model="formData.lastOnlineQueryTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
// 字典增加如下代码
const BannedStatusOptions = ref([])
const OnlineStatusOptions = ref([])

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
             roleOnlineStatus: '',
            scriptOnlineStatus: '',
            bannedStatus: '',
            talismanDiff: undefined,
            ingotDiff: undefined,
            lastSyncQueryTime: new Date(),
            lastSyncUpdateTime: new Date(),
            scriptLastOnlineTime: new Date(),
            roleLastOnlineTime: new Date(),
            roleGameId: '',
            serverZoneId: '',
            lastIngotTradeTime: new Date(),
            lastTalismanTradeTime: new Date(),
            onlineTalismanTotal: undefined,
            onlineIngotTotal: undefined,
            lastOnlineQueryTime: new Date(),
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
    BannedStatusOptions.value = await getDictFunc('BannedStatus')
    OnlineStatusOptions.value = await getDictFunc('OnlineStatus')

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
