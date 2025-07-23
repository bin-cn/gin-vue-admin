
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="登录码" prop="loginCode">
  <el-input v-model="searchInfo.loginCode" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="机器编号" prop="machineNoName">
  <el-input v-model="searchInfo.machineNoName" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="区服" prop="gameServerName">
  <el-input v-model="searchInfo.gameServerName" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="区服id" prop="gameServerId">
  <el-input class="w-40" v-model.number="searchInfo.startGameServerId" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endGameServerId" placeholder="最大值" />
</el-form-item>
            
            <el-form-item label="游戏角色ID" prop="roleGameId">
  <el-input v-model.number="searchInfo.roleGameId" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="ID名字" prop="iDName">
  <el-input v-model="searchInfo.iDName" placeholder="搜索条件" />
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="smartcreate_GameUserAuthCode" />
            <ExportExcel  template-id="smartcreate_GameUserAuthCode" filterDeleted/>
            <ImportExcel  template-id="smartcreate_GameUserAuthCode" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="登录码" prop="loginCode" width="120" />

            <el-table-column sortable align="left" label="机器编号" prop="machineNoName" width="120" />

            <el-table-column sortable align="left" label="使用人" prop="assignerName" width="120" />

            <el-table-column sortable align="left" label="账号" prop="account" width="120" />

            <el-table-column align="left" label="密码" prop="password" width="120" />

            <el-table-column sortable align="left" label="区服" prop="gameServerName" width="120" />

            <el-table-column sortable align="left" label="区服id" prop="gameServerId" width="120" />

            <el-table-column align="left" label="游戏角色名字" prop="roleGameName" width="120" />

            <el-table-column sortable align="left" label="游戏角色ID" prop="roleGameId" width="120" />

            <el-table-column sortable align="left" label="开区时间" prop="serverOpenTime" width="120" />

            <el-table-column sortable align="left" label="进区时间" prop="enterServerTime" width="120" />

            <el-table-column sortable align="left" label="账号状态" prop="accountStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.accountStatus,AccountStatusOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="备注" prop="remark" width="120" />

            <el-table-column sortable align="left" label="ID名字" prop="iDName" width="120" />

            <el-table-column sortable align="left" label="身份证号码" prop="iDCardNumber" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateGameUserAuthCodeFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="登录码">
    {{ detailForm.loginCode }}
</el-descriptions-item>
                    <el-descriptions-item label="机器编号">
    {{ detailForm.machineNoName }}
</el-descriptions-item>
                    <el-descriptions-item label="使用人">
    {{ detailForm.assignerName }}
</el-descriptions-item>
                    <el-descriptions-item label="账号">
    {{ detailForm.account }}
</el-descriptions-item>
                    <el-descriptions-item label="密码">
    {{ detailForm.password }}
</el-descriptions-item>
                    <el-descriptions-item label="区服">
    {{ detailForm.gameServerName }}
</el-descriptions-item>
                    <el-descriptions-item label="区服id">
    {{ detailForm.gameServerId }}
</el-descriptions-item>
                    <el-descriptions-item label="游戏角色名字">
    {{ detailForm.roleGameName }}
</el-descriptions-item>
                    <el-descriptions-item label="游戏角色ID">
    {{ detailForm.roleGameId }}
</el-descriptions-item>
                    <el-descriptions-item label="开区时间">
    {{ detailForm.serverOpenTime }}
</el-descriptions-item>
                    <el-descriptions-item label="进区时间">
    {{ detailForm.enterServerTime }}
</el-descriptions-item>
                    <el-descriptions-item label="账号状态">
    {{ detailForm.accountStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="备注">
    {{ detailForm.remark }}
</el-descriptions-item>
                    <el-descriptions-item label="ID名字">
    {{ detailForm.iDName }}
</el-descriptions-item>
                    <el-descriptions-item label="身份证号码">
    {{ detailForm.iDCardNumber }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createGameUserAuthCode,
  deleteGameUserAuthCode,
  deleteGameUserAuthCodeByIds,
  updateGameUserAuthCode,
  findGameUserAuthCode,
  getGameUserAuthCodeList
} from '@/api/smartcreate/uGameUserAuthCode'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'GameUserAuthCode'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               machineNoName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               assignerName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               account : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               gameServerName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               gameServerId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               roleGameName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               roleGameId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               iDName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               iDCardNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"CreatedAt",
    ID:"ID",
            machineNoName: 'machine_no_name',
            assignerName: 'assigner_name',
            account: 'account',
            gameServerName: 'game_server_name',
            gameServerId: 'game_server_id',
            roleGameId: 'role_game_id',
            serverOpenTime: 'server_open_time',
            enterServerTime: 'enter_server_time',
            accountStatus: 'account_status',
            iDName: 'i_d_name',
            iDCardNumber: 'i_d_card_number',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getGameUserAuthCodeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    AccountStatusOptions.value = await getDictFunc('AccountStatus')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteGameUserAuthCodeFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteGameUserAuthCodeByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateGameUserAuthCodeFunc = async(row) => {
    const res = await findGameUserAuthCode({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteGameUserAuthCodeFunc = async (row) => {
    const res = await deleteGameUserAuthCode({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findGameUserAuthCode({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
