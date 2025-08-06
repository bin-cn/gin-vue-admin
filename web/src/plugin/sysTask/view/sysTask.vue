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
        <el-form-item
          label="CreatedTime"
          prop="createdAt"
        >
          <template #label>
            <span>
              CreatedTime
              <el-tooltip content="The search range is start date (inclusive) to end date (exclusive)">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="StareDate"
            :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"
          />
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="EndDate"
            :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"
          />
        </el-form-item>
        <el-form-item label="InvokeTarget">
          <el-input
              clearable
              v-model="searchInfo.invokeTarget"
              placeholder="invokeTarget"
          />
        </el-form-item>
        <el-form-item label="分组">
          <el-input
              clearable
              v-model="searchInfo.group"
              placeholder="分组名"
          />
        </el-form-item>
        <el-form-item label="Cron">
          <el-input
              clearable
              v-model="searchInfo.cronExpression"
              placeholder="Cron"
          />
        </el-form-item>
        <el-form-item label="描述">
          <el-input
              clearable
              v-model="searchInfo.remark"
              placeholder="Remark"
          />
        </el-form-item>
        <el-form-item label="状态">
          <el-select
              v-model="searchInfo.status"
              placeholder="请选择"
              style="width: 130px"
              clearable
          >
            <el-option
                v-for="item in statusOption"
                :key="item.label"
                width="80"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
          <el-form-item label="任务类型">
            <el-select
                v-model="searchInfo.tag"
                placeholder="请选择"
                style="width: 130px"
                clearable
            >
              <el-option
                  v-for="item in taskTypeOption"
                  :key="item.label"
                  width="80"
                  :label="item.label"
                  :value="item.value"
              />
            </el-select>
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
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >newAdd</el-button>
      </div>
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
          label="CreatedAt"
          width="160"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="invokeTarget"
          prop="invokeTarget"
          width="120"
        />
        <el-table-column
          align="left"
          label="group"
          prop="taskGroup"
          width="80"
        />
        <el-table-column
          align="left"
          label="cron"
          prop="cronExpression"
          width="120"
        />
        <el-table-column
          align="left"
          label="args"
          prop="args"
          width="80"
        />
        <el-table-column
          align="left"
          label="remark"
          prop="remark"
          width="220"
        />
        <el-table-column
          align="left"
          label="status"
          prop="status"
          width="120"
        >
          <template #default="scope">{{ showDictLabel(sysTaskStatus, scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="action" fixed="right" min-width="140">
          <template #default="scope">
            <el-button type="primary" icon="edit" plain class="table-button" @click="updateSysTaskFunc(scope.row)">Update</el-button><br/>
            <el-button
                type="warning"
                plain
                icon="delete"
                @click="remove(scope.row)"
            >Delete</el-button>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="operate"
          fixed="right"
          min-width="160"
        >
          <template #default="scope">
            <el-button
              type="primary"
              plain
              @click="enable(scope.row)"
            >Start</el-button>
            <el-button
              type="danger"
              plain
              @click="disable(scope.row)"
            >Stop</el-button><br>
            <el-button
              type="success"
              plain
              @click="runOne(scope.row)"
            >RunOne</el-button>

            <!--              <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">-->
            <!--                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>-->
            <!--                CheckDetails-->
            <!--            </el-button>-->
            <!--            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">Delete</el-button>-->
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
      <el-scrollbar height="600px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="120px"
        >
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
            label="任务分组:"
            prop="taskGroup"
          >
            <el-input
              v-model="formData.taskGroup"
              :clearable="true"
              placeholder="请输入任务分组"
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
            label="invokeTarget:"
            prop="invokeTarget"
          >
            <el-input
              v-model="formData.invokeTarget"
              :clearable="true"
              placeholder="请输入invokeTarget"
            />
          </el-form-item>
          <el-form-item
            label="args:"
            prop="args"
          >
            <el-input
              v-model="formData.args"
              :clearable="true"
              placeholder="请输入args，或http的Body"
            />
          </el-form-item>
          <el-form-item
              label="httpMethod方法:"
              prop="httpMethod"
          >
            <el-input
                v-model="formData.httpMethod"
                :clearable="true"
                placeholder="类型是http请求，请输入method，目前只支持GET、POST"
            />
          </el-form-item>
          <el-form-item
            label="timeout"
            prop="timeout"
          >
            <el-input
              v-model.number="formData.timeout"
              :clearable="true"
              placeholder="请输入任务执行超时时间(单位秒),0不限制"
            />
          </el-form-item>
          <el-form-item
            label="multi:"
            prop="multi"
          >
            <el-input
              v-model="formData.multi"
              :clearable="true"
              placeholder="是否允许多实例运行"
            />
          </el-form-item>
          <el-form-item
            label="retryTimes:"
            prop="retryTimes"
          >
            <el-input
              v-model="formData.retryTimes"
              :clearable="true"
              placeholder="重试次数"
            />
          </el-form-item>
          <el-form-item
            label="retryInterval:"
            prop="retryInterval"
          >
            <el-input
              v-model="formData.retryInterval"
              :clearable="true"
              placeholder="请输入重试间隔时间"
            />
          </el-form-item>
          <el-form-item
              label="Tag"
              prop="tag"
          >
            <el-select
                v-model="formData.tag"
                placeholder="请选择"
                :clearable="true"
                style="width:100%"
            >
              <el-option
                  v-for="item in taskTypeOption"
                  :key="item.value"
                  :label="`${item.label}`"
                  :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            label="remark:"
            prop="remark"
          >
            <el-input
              v-model="formData.remark"
              :clearable="true"
              placeholder="请输入remark"
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
          <el-descriptions-item label="任务名称">
            {{ formData.taskName }}
          </el-descriptions-item>
          <el-descriptions-item label="任务分组">
            {{ formData.taskGroup }}
          </el-descriptions-item>
          <el-descriptions-item label="cron_expression">
            {{ formData.cronExpression }}
          </el-descriptions-item>
          <el-descriptions-item label="invokeTarget">
            {{ formData.invokeTarget }}
          </el-descriptions-item>
          <el-descriptions-item label="args">
            {{ formData.args }}
          </el-descriptions-item>
          <el-descriptions-item label="任务执行超时时间(单位秒),0不限制">
            {{ formData.timeout }}
          </el-descriptions-item>
          <el-descriptions-item label="是否允许多实例运行">
            {{ formatBoolean(formData.multi) }}
          </el-descriptions-item>
          <el-descriptions-item label="重试次数">
            {{ formatBoolean(formData.retryTimes) }}
          </el-descriptions-item>
          <el-descriptions-item label="重试间隔时间">
            {{ formData.retryInterval }}
          </el-descriptions-item>
          <el-descriptions-item label="tag">
            {{ formData.tag }}
          </el-descriptions-item>
          <el-descriptions-item label="remark">
            {{ formData.remark }}
          </el-descriptions-item>
          <el-descriptions-item label="状态 1:正常 0:停止">
            {{ formatBoolean(formData.status) }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSysTask,
  deleteSysTask,
  deleteSysTaskByIds,
  updateSysTask,
  findSysTask,
  getSysTaskList,
  runOneSysTask,
  enableSysTask,
  disableSysTask,
  removeSysTask
} from '@/plugin/sysTask/api/sysTask'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { getDict, showDictLabel } from '@/utils/dictionary'

defineOptions({
  name: 'SysTask'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  taskName: '',
  taskGroup: '',
  cronExpression: '',
  invokeTarget: '',
  args: '',
  timeout: 0,
  multi: 1,
  retryTimes: 0,
  retryInterval: 0,
  tag: 'internalFuncTask',
  remark: '',
  status: 1,
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
const sysTaskType = ref([])
const statusOption = ref([])
const taskTypeOption = ref([])


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
    if (searchInfo.value.multi === '') {
      searchInfo.value.multi = null
    }
    if (searchInfo.value.retryTimes === '') {
      searchInfo.value.retryTimes = null
    }
    if (searchInfo.value.status === '') {
      searchInfo.value.status = null
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
  const table = await getSysTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  sysTaskStatus.value = await getDict('sysTaskStatus')
  sysTaskType.value = await getDict('sysTaskType')
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
  statusOption.value = await getDictFunc('sysTaskStatus')
  taskTypeOption.value = await  getDictFunc('sysTaskType')
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
    deleteSysTaskFunc(row)
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
  const res = await deleteSysTaskByIds({ IDs })
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
const updateSysTaskFunc = async(row) => {
  const res = await findSysTask({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysTask
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteSysTaskFunc = async(row) => {
  const res = await removeSysTask({ ID: row.ID })
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
  const res = await findSysTask({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.resysTask
    openDetailShow()
  }
}

// 启动
const enable = async(row) => {
  const res = await enableSysTask({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '启动成功'
    })
    getTableData()
  }
}

// 暂停任务
const disable = async(row) => {
  const res = await disableSysTask({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '暂停任务成功'
    })
    getTableData()
  }
}

// 运行一次
const runOne = async(row) => {
  const res = await runOneSysTask({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '运行一次成功'
    })
  }
}

const remove = async(row) => {
  ElMessageBox.confirm('You sure you want to delete it?', 'Hint', {
    confirmButtonText: 'Ok',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    deleteSysTaskFunc(row)
    // const res = removeSysTask({ ID: row.ID })
    // if (res.code === 0) {
    //   ElMessage({
    //     type: 'success',
    //     message: res.msg
    //   })
    //   setTimeout(() => {
    //     getTableData()
    //   }, 1000)
    // }
  })
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    taskName: '',
    taskGroup: '',
    cronExpression: '',
    invokeTarget: '',
    args: '',
    timeout: 0,
    multi: 1,
    retryTimes: 0,
    retryInterval: 0,
    tag: '',
    remark: '',
    status: 1,
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
    taskName: '',
    taskGroup: '',
    cronExpression: '',
    invokeTarget: '',
    args: '',
    timeout: 0,
    multi: 1,
    retryTimes: 0,
    retryInterval: 0,
    tag: 'internalFuncTask',
    remark: '',
    status: 1,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
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
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>
