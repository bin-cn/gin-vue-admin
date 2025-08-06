<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="invokeTarget">
          <el-input
            v-model="searchInfo.invokeTarget"
            placeholder="invokeTarget Name"
            clearable
          />
        </el-form-item>
        <el-form-item label="cronExpression">
          <el-input
            v-model="searchInfo.cronExpression"
            placeholder="cronExpression Name"
            clearable
          />
        </el-form-item>
        <el-form-item label="result">
          <el-input
            v-model="searchInfo.result"
            placeholder="result"
            clearable
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >Search</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="target"
          prop="invokeTarget"
          width="130"
        />
        <el-table-column
          align="left"
          label="cron"
          prop="cronExpression"
          width="120"
        />
        <el-table-column
          align="left"
          label="startTime"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.startTime) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="endTime"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.endTime) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="status"
          prop="status"
          width="80"
        >
          <template #default="scope">{{ showDictLabel(sysTaskStatus, scope.row.status) }}</template>
        </el-table-column>
        <el-table-column
          align="center"
          label="result"
          prop="result"
          width="220"
        >
          <template #default="scope">
            <el-tooltip
                :content="scope.row.result"
                placement="top"
            >
              <div>{{ truncateLength(scope.row.result) }}</div>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="totalSecond"
          prop="totalTime"
          width="110"
        >
          <template #default="scope">{{ scope.row.totalTime }}s</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="operate"
          fixed="right"
          min-width="140"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              CheckDetails
            </el-button>
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type==='create'?'Add':'Update'"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="80px"
        >
          <el-form-item
            label="任务id:"
            prop="taskId"
          >
            <el-input
              v-model.number="formData.taskId"
              :clearable="true"
              placeholder="请输入任务id"
            />
          </el-form-item>
          <el-form-item
            label="任务名称:"
            prop="taskName"
          >
            <el-input
              v-model="formData.taskName"
              :clearable="true"
              placeholder="请输入任务名称"
            />
          </el-form-item>
          <el-form-item
            label="cron_expression:"
            prop="cronExpression"
          >
            <el-input
              v-model="formData.cronExpression"
              :clearable="true"
              placeholder="请输入cron_expression"
            />
          </el-form-item>
          <el-form-item
            label="invoke_target:"
            prop="invokeTarget"
          >
            <el-input
              v-model="formData.invokeTarget"
              :clearable="true"
              placeholder="请输入invoke_target"
            />
          </el-form-item>
          <el-form-item
            label="timeout:"
            prop="timeout"
          >
            <el-input
              v-model.number="formData.timeout"
              :clearable="true"
              placeholder="请输入timeout"
            />
          </el-form-item>
          <el-form-item
            label="retryTimes:"
            prop="retryTimes"
          >
            <el-switch
              v-model="formData.retryTimes"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="startTime:"
            prop="startTime"
          >
            <el-date-picker
              v-model="formData.startTime"
              type="date"
              style="width:100%"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item
            label="endTime:"
            prop="endTime"
          >
            <el-date-picker
              v-model="formData.endTime"
              type="date"
              style="width:100%"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item
            label="status:"
            prop="status"
          >
            <el-switch
              v-model="formData.status"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
              clearable
            />
          </el-form-item>
          <el-form-item
            label="result:"
            prop="result"
          >
            <el-input
              v-model="formData.result"
              :clearable="true"
              placeholder="请输入result"
            />
          </el-form-item>
          <el-form-item
            label="totalTime:"
            prop="totalTime"
          >
            <el-input
              v-model.number="formData.totalTime"
              :clearable="true"
              placeholder="请输入totalTime"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >Ok</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="Check details"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="任务id">
            {{ formData.taskId }}
          </el-descriptions-item>
          <el-descriptions-item label="任务名称">
            {{ formData.taskName }}
          </el-descriptions-item>
          <el-descriptions-item label="cron_expression">
            {{ formData.cronExpression }}
          </el-descriptions-item>
          <el-descriptions-item label="invoke_target">
            {{ formData.invokeTarget }}
          </el-descriptions-item>
          <el-descriptions-item label="timeout">
            {{ formData.timeout }}
          </el-descriptions-item>
          <el-descriptions-item label="retryTimes">
            {{ formatBoolean(formData.retryTimes) }}
          </el-descriptions-item>
          <el-descriptions-item label="startTime">
            {{ formatDate(formData.startTime) }}
          </el-descriptions-item>
          <el-descriptions-item label="endTime">
            {{ formatDate(formData.endTime) }}
          </el-descriptions-item>
          <el-descriptions-item label="status">
            {{ formatBoolean(formData.status) }}
          </el-descriptions-item>
          <el-descriptions-item label="result">
            {{ formData.result }}
          </el-descriptions-item>
          <el-descriptions-item label="totalTime">
            {{ formData.totalTime }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSysTaskLog,
  deleteSysTaskLog,
  deleteSysTaskLogByIds,
  updateSysTaskLog,
  findSysTaskLog,
  getSysTaskLogList
} from '@/plugin/sysTask/api/sysTaskLog'

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean } from '@/utils/format'
import { truncateLength } from '@/plugin/sysTask/api/my_format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getDict, showDictLabel } from '@/utils/dictionary'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'SysTaskLog'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  taskId: 0,
  taskName: '',
  cronExpression: '',
  invokeTarget: '',
  timeout: 0,
  retryTimes: 1,
  startTime: new Date(),
  endTime: new Date(),
  status: -1,
  result: '',
  totalTime: 0,
})

// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
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
const sysTaskStatus = ref([])

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
    pageSize.value = 10
    if (searchInfo.value.retryTimes === '') {
      searchInfo.value.retryTimes = null
    }
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
  const table = await getSysTaskLogList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  sysTaskStatus.value = await getDict('sysTaskStatus')
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
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
    deleteSysTaskLogFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
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
  const res = await deleteSysTaskLogByIds({ IDs })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === IDs.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateSysTaskLogFunc = async(row) => {
  const res = await findSysTaskLog({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysTaskLog
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteSysTaskLogFunc = async(row) => {
  const res = await deleteSysTaskLog({ ID: row.ID })
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

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findSysTaskLog({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.resysTaskLog
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
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
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
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
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>
