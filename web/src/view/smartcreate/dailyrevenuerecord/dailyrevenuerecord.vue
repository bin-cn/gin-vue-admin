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
            <ExportTemplate  template-id="smartcreate_DailyRevenueRecord" />
            <ExportExcel  template-id="smartcreate_DailyRevenueRecord" filterDeleted/>
            <ImportExcel  template-id="smartcreate_DailyRevenueRecord" @on-success="getTableData" />
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
        
            <el-table-column sortable align="left" label="统计日期" prop="statisticDate" width="120" />

            <el-table-column sortable align="left" label="用户ID" prop="userId" width="120" />

            <el-table-column sortable align="left" label="用户昵称" prop="userNickname" width="120" />

            <el-table-column sortable align="left" label="区服名字" prop="serverName" width="120" />

            <el-table-column sortable align="left" label="区服ID" prop="serverZoneId" width="120" />

            <el-table-column sortable align="left" label="主区服ID" prop="mainServerZoneId" width="120" />

            <el-table-column sortable align="left" label="统计类型" prop="statisticType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.statisticType,MaterialTypeOptions) }}
    </template>
</el-table-column>
            <el-table-column sortable align="left" label="产出金额" prop="amount" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateDailyRevenueRecordFunc(scope.row)">编辑</el-button>
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="统计日期">
    {{ detailForm.statisticDate }}
</el-descriptions-item>
                    <el-descriptions-item label="用户ID">
    {{ detailForm.userId }}
</el-descriptions-item>
                    <el-descriptions-item label="用户昵称">
    {{ detailForm.userNickname }}
</el-descriptions-item>
                    <el-descriptions-item label="区服名字">
    {{ detailForm.serverName }}
</el-descriptions-item>
                    <el-descriptions-item label="区服ID">
    {{ detailForm.serverZoneId }}
</el-descriptions-item>
                    <el-descriptions-item label="主区服ID">
    {{ detailForm.mainServerZoneId }}
</el-descriptions-item>
                    <el-descriptions-item label="统计类型">
    {{ detailForm.statisticType }}
</el-descriptions-item>
                    <el-descriptions-item label="产出金额">
    {{ detailForm.amount }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createDailyRevenueRecord,
  deleteDailyRevenueRecord,
  deleteDailyRevenueRecordByIds,
  updateDailyRevenueRecord,
  findDailyRevenueRecord,
  getDailyRevenueRecordList
} from '@/api/smartcreate/dailyrevenuerecord'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, } from 'vue'
import { useAppStore } from "@/pinia"


// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'DailyRevenueRecord'
})



// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               userId : [{
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
               userNickname : [{
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
               serverName : [{
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
               serverZoneId : [{
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
               mainServerZoneId : [{
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
               statisticType : [{
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
               amount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
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
            statisticDate: 'statistic_date',
            userId: 'user_id',
            userNickname: 'user_nickname',
            serverName: 'server_name',
            serverZoneId: 'server_zone_id',
            mainServerZoneId: 'main_server_zone_id',
            statisticType: 'statistic_type',
            amount: 'amount',
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
  const table = await getDailyRevenueRecordList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    MaterialTypeOptions.value = await getDictFunc('MaterialType')
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
            deleteDailyRevenueRecordFunc(row)
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
      const res = await deleteDailyRevenueRecordByIds({ IDs })
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
const updateDailyRevenueRecordFunc = async(row) => {
    const res = await findDailyRevenueRecord({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDailyRevenueRecordFunc = async (row) => {
    const res = await deleteDailyRevenueRecord({ ID: row.ID })
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
        statisticDate: '',
        userId: '',
        userNickname: '',
        serverName: '',
        serverZoneId: '',
        mainServerZoneId: '',
        statisticType: '',
        amount: 0,
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
  const res = await findDailyRevenueRecord({ ID: row.ID })
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
