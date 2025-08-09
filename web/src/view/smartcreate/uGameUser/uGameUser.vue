
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
      
            <el-form-item label="用户昵称" prop="nickName">
  <el-select v-model="searchInfo.nickName" filterable placeholder="请选择用户昵称" :clearable="true">
    <el-option v-for="(item,key) in dataSource.nickName" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="登录码" prop="loginCode">
  <el-input v-model="searchInfo.loginCode" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="用户ID" prop="userId">
  <el-select v-model="searchInfo.userId" filterable placeholder="请选择用户ID" :clearable="true">
    <el-option v-for="(item,key) in dataSource.userId" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            
            <el-form-item label="未绑定元宝数量" prop="unBoundIngotQuantity">
  <el-input class="w-40" v-model.number="searchInfo.startUnBoundIngotQuantity" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endUnBoundIngotQuantity" placeholder="最大值" />
</el-form-item>
            
            <el-form-item label="绑定元宝数量" prop="boundIngotQuantity">
  <el-input class="w-40" v-model.number="searchInfo.startBoundIngotQuantity" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endBoundIngotQuantity" placeholder="最大值" />
</el-form-item>
            
            <el-form-item label="元宝总数" prop="totalIngotQuantity">
  <el-input class="w-40" v-model.number="searchInfo.startTotalIngotQuantity" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endTotalIngotQuantity" placeholder="最大值" />
</el-form-item>
            
            <el-form-item label="未绑定灵符" prop="unBoundTalisman">
  <el-input class="w-40" v-model.number="searchInfo.startUnBoundTalisman" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endUnBoundTalisman" placeholder="最大值" />
</el-form-item>
            
            <el-form-item label="绑定灵符" prop="boundTalisman">
  <el-input class="w-40" v-model.number="searchInfo.startBoundTalisman" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endBoundTalisman" placeholder="最大值" />
</el-form-item>
            
            <el-form-item label="灵符总数" prop="totalTalisman">
  <el-input class="w-40" v-model.number="searchInfo.startTotalTalisman" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endTotalTalisman" placeholder="最大值" />
</el-form-item>

  <el-form-item label="区服名称" prop="gameServerName">
  <el-input v-model="searchInfo.gameServerName" placeholder="搜索条件" />
</el-form-item>

<el-form-item label="角色在线状态" prop="roleOnlineStatus">
  <el-select v-model="searchInfo.roleOnlineStatus" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.roleOnlineStatus=undefined}">
    <el-option v-for="(item,key) in OnlineStatusOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
    

<el-form-item label="脚本在线状态" prop="scriptOnlineStatus">
  <el-select v-model="searchInfo.scriptOnlineStatus" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.scriptOnlineStatus=undefined}">
    <el-option v-for="(item,key) in OnlineStatusOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
    

<el-form-item label="封号状态" prop="bannedStatus">
  <el-select v-model="searchInfo.bannedStatus" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.bannedStatus=undefined}">
    <el-option v-for="(item,key) in BannedStatusOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
    

<el-form-item label="灵符变动差异" prop="talismanDiff">
  <el-input class="w-40" v-model.number="searchInfo.startTalismanDiff" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endTalismanDiff" placeholder="最大值" />
</el-form-item>
    

<el-form-item label="元宝变动差异" prop="ingotDiff">
  <el-input class="w-40" v-model.number="searchInfo.startIngotDiff" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endIngotDiff" placeholder="最大值" />
</el-form-item>
    

<el-form-item label="最后一次同步查询时间" prop="lastSyncQueryTime">
  <template #label>
    <span>
      最后一次同步查询时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="w-[380px]" v-model="searchInfo.lastSyncQueryTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
    

<el-form-item label="最后一次同步更新时间" prop="lastSyncUpdateTime">
  <template #label>
    <span>
      最后一次同步更新时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="w-[380px]" v-model="searchInfo.lastSyncUpdateTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
    

<el-form-item label="脚本最后在线时间" prop="scriptLastOnlineTime">
  <template #label>
    <span>
      脚本最后在线时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="w-[380px]" v-model="searchInfo.scriptLastOnlineTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
    

<el-form-item label="角色最后在线时间" prop="roleLastOnlineTime">
  <template #label>
    <span>
      角色最后在线时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="w-[380px]" v-model="searchInfo.roleLastOnlineTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
    

<el-form-item label="游戏角色ID" prop="roleGameId">
  <el-input v-model="searchInfo.roleGameId" placeholder="搜索条件" />
</el-form-item>
    

<el-form-item label="原始区服ID" prop="serverZoneId">
  <el-input v-model="searchInfo.serverZoneId" placeholder="搜索条件" />
</el-form-item>
    

<el-form-item label="最后一次交易元宝变动时间" prop="lastIngotTradeTime">
  <template #label>
    <span>
      最后一次交易元宝变动时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="w-[380px]" v-model="searchInfo.lastIngotTradeTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
    

<el-form-item label="最后一次交易灵符时间" prop="lastTalismanTradeTime">
  <template #label>
    <span>
      最后一次交易灵符时间
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="w-[380px]" v-model="searchInfo.lastTalismanTradeTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
    

<el-form-item label="线上灵符总数" prop="onlineTalismanTotal">
  <el-input v-model.number="searchInfo.onlineTalismanTotal" placeholder="搜索条件" />
</el-form-item>
    

<el-form-item label="线上元宝总数" prop="onlineIngotTotal">
  <el-input v-model.number="searchInfo.onlineIngotTotal" placeholder="搜索条件" />
</el-form-item>
    
<el-form-item label="区服ID" prop="gameServerId">
  <el-input class="w-40" v-model.number="searchInfo.startGameServerId" placeholder="最小值" />
  —
  <el-input class="w-40" v-model.number="searchInfo.endGameServerId" placeholder="最大值" />
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
            <ExportTemplate  template-id="smartcreate_GameUser" />
            <ExportExcel  template-id="smartcreate_GameUser" filterDeleted/>
            <ImportExcel  template-id="smartcreate_GameUser" @on-success="getTableData" />
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
        
        <el-table-column sortable align="left" label="区服名称" prop="gameServerName" width="120" />

        <el-table-column sortable align="left" label="区服ID" prop="gameServerId" width="120" />


        <el-table-column sortable align="left" label="用户昵称" prop="nickName" width="120">
                <template #default="scope">
                    <span>{{ filterDataSource(dataSource.nickName,scope.row.nickName) }}</span>
                </template>
        </el-table-column>


            <el-table-column align="left" label="登录码" prop="loginCode" width="120" />

            <el-table-column align="left" label="游戏角色名称" prop="roleGameName" width="120" />

            <el-table-column sortable align="left" label="游戏角色等级" prop="roleGameLevel" width="120" />

            <el-table-column align="left" label="用户ID" prop="userId" width="120">
             <template #default="scope">
        <span>{{ filterDataSource(dataSource.userId,scope.row.userId) }}</span>
    </template>
</el-table-column>

            <el-table-column sortable align="left" label="角色在线状态" prop="roleOnlineStatus" width="120">
                <template #default="scope">
                {{ filterDict(scope.row.roleOnlineStatus,OnlineStatusOptions) }}
                </template>
          </el-table-column>


       <el-table-column sortable align="left" label="脚本在线状态" prop="scriptOnlineStatus" width="120">
          <template #default="scope">
          {{ filterDict(scope.row.scriptOnlineStatus,OnlineStatusOptions) }}
          </template>
      </el-table-column>
       <el-table-column sortable align="left" label="封号状态" prop="bannedStatus" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.bannedStatus,BannedStatusOptions) }}
    </template>
</el-table-column>
       <el-table-column sortable align="left" label="灵符变动差异" prop="talismanDiff" width="120" />

       <el-table-column sortable align="left" label="元宝变动差异" prop="ingotDiff" width="120" />

       <el-table-column sortable align="left" label="最后一次同步查询时间" prop="lastSyncQueryTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastSyncQueryTime) }}</template>
</el-table-column>
       <el-table-column sortable align="left" label="最后一次同步更新时间" prop="lastSyncUpdateTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastSyncUpdateTime) }}</template>
</el-table-column>
       <el-table-column sortable align="left" label="脚本最后在线时间" prop="scriptLastOnlineTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.scriptLastOnlineTime) }}</template>
</el-table-column>
       <el-table-column sortable align="left" label="角色最后在线时间" prop="roleLastOnlineTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.roleLastOnlineTime) }}</template>
</el-table-column>
       <el-table-column align="left" label="游戏角色ID" prop="roleGameId" width="120" />

       <el-table-column sortable align="left" label="原始区服ID" prop="serverZoneId" width="120" />

       <el-table-column sortable align="left" label="最后一次交易元宝变动时间" prop="lastIngotTradeTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastIngotTradeTime) }}</template>
</el-table-column>
       <el-table-column sortable align="left" label="最后一次交易灵符时间" prop="lastTalismanTradeTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastTalismanTradeTime) }}</template>
</el-table-column>
       <el-table-column sortable align="left" label="线上灵符总数" prop="onlineTalismanTotal" width="120" />

       <el-table-column sortable align="left" label="线上元宝总数" prop="onlineIngotTotal" width="120" />

       <el-table-column align="left" label="最后一次线上查询时间" prop="lastOnlineQueryTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.lastOnlineQueryTime) }}</template>
</el-table-column>   
              
   
            <el-table-column sortable align="left" label="未绑定元宝数量" prop="unBoundIngotQuantity" width="120" />

            <el-table-column sortable align="left" label="绑定元宝数量" prop="boundIngotQuantity" width="120" />

            <el-table-column sortable align="left" label="元宝总数" prop="totalIngotQuantity" width="120" />

            <el-table-column sortable align="left" label="未绑定灵符" prop="unBoundTalisman" width="120" />

            <el-table-column sortable align="left" label="绑定灵符" prop="boundTalisman" width="120" />

            <el-table-column sortable align="left" label="灵符总数" prop="totalTalisman" width="120" />
            
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateGameUserFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="用户昵称:" prop="nickName">
    <el-select v-model="formData.nickName" placeholder="请选择用户昵称" filterable style="width:100%" :clearable="true">
        <el-option v-for="(item,key) in dataSource.nickName" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>

   <el-form-item label="区服名称:" prop="gameServerName">
    <el-input v-model="formData.gameServerName" :clearable="true" placeholder="请输入区服名称" />
</el-form-item>
     <el-form-item label="区服ID:" prop="gameServerId">
    <el-input v-model.number="formData.gameServerId" :clearable="true" placeholder="请输入区服ID" />
</el-form-item>

<el-form-item label="登录码:" prop="loginCode">
    <el-input v-model="formData.loginCode" :clearable="true" placeholder="请输入登录码" />
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



          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="用户昵称">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.nickName,detailForm.nickName) }}</span>
    </template>
</el-descriptions-item>

  <el-descriptions-item label="区服名称">
    {{ detailForm.gameServerName }}
</el-descriptions-item>
    <el-descriptions-item label="区服ID">
    {{ detailForm.gameServerId }}
</el-descriptions-item>


<el-descriptions-item label="登录码">
    {{ detailForm.loginCode }}
</el-descriptions-item>
                    <el-descriptions-item label="游戏角色名称">
    {{ detailForm.roleGameName }}
</el-descriptions-item>
                    <el-descriptions-item label="游戏角色等级">
    {{ detailForm.roleGameLevel }}
</el-descriptions-item>
                    <el-descriptions-item label="用户ID">
    <template #default="scope">
        <span>{{ filterDataSource(dataSource.userId,detailForm.userId) }}</span>
    </template>
</el-descriptions-item>
                    <el-descriptions-item label="未绑定元宝数量">
    {{ detailForm.unBoundIngotQuantity }}
</el-descriptions-item>
                    <el-descriptions-item label="绑定元宝数量">
    {{ detailForm.boundIngotQuantity }}
</el-descriptions-item>
                    <el-descriptions-item label="元宝总数">
    {{ detailForm.totalIngotQuantity }}
</el-descriptions-item>
                    <el-descriptions-item label="未绑定灵符">
    {{ detailForm.unBoundTalisman }}
</el-descriptions-item>
                    <el-descriptions-item label="绑定灵符">
    {{ detailForm.boundTalisman }}
</el-descriptions-item>
                    <el-descriptions-item label="灵符总数">
    {{ detailForm.totalTalisman }}
</el-descriptions-item>

// 查看抽屉中增加如下代码
    <el-descriptions-item label="角色在线状态">
    {{ detailForm.roleOnlineStatus }}
</el-descriptions-item>
    <el-descriptions-item label="脚本在线状态">
    {{ detailForm.scriptOnlineStatus }}
</el-descriptions-item>
    <el-descriptions-item label="封号状态">
    {{ detailForm.bannedStatus }}
</el-descriptions-item>
    <el-descriptions-item label="灵符变动差异">
    {{ detailForm.talismanDiff }}
</el-descriptions-item>
    <el-descriptions-item label="元宝变动差异">
    {{ detailForm.ingotDiff }}
</el-descriptions-item>
    <el-descriptions-item label="最后一次同步查询时间">
    {{ detailForm.lastSyncQueryTime }}
</el-descriptions-item>
    <el-descriptions-item label="最后一次同步更新时间">
    {{ detailForm.lastSyncUpdateTime }}
</el-descriptions-item>
    <el-descriptions-item label="脚本最后在线时间">
    {{ detailForm.scriptLastOnlineTime }}
</el-descriptions-item>
    <el-descriptions-item label="角色最后在线时间">
    {{ detailForm.roleLastOnlineTime }}
</el-descriptions-item>
    <el-descriptions-item label="游戏角色ID">
    {{ detailForm.roleGameId }}
</el-descriptions-item>
    <el-descriptions-item label="原始区服ID">
    {{ detailForm.serverZoneId }}
</el-descriptions-item>
    <el-descriptions-item label="最后一次交易元宝变动时间">
    {{ detailForm.lastIngotTradeTime }}
</el-descriptions-item>
    <el-descriptions-item label="最后一次交易灵符时间">
    {{ detailForm.lastTalismanTradeTime }}
</el-descriptions-item>
    <el-descriptions-item label="线上灵符总数">
    {{ detailForm.onlineTalismanTotal }}
</el-descriptions-item>
    <el-descriptions-item label="线上元宝总数">
    {{ detailForm.onlineIngotTotal }}
</el-descriptions-item>
    <el-descriptions-item label="最后一次线上查询时间">
    {{ detailForm.lastOnlineQueryTime }}
</el-descriptions-item>

            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
    getGameUserDataSource,
  createGameUser,
  deleteGameUser,
  deleteGameUserByIds,
  updateGameUser,
  findGameUser,
  getGameUserList
} from '@/api/smartcreate/uGameUser'

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
    name: 'GameUser'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)
// 字典增加如下代码
const BannedStatusOptions = ref([])
const OnlineStatusOptions = ref([])


// 自动化生成的字典（可能为空）以及字段
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
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getGameUserDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               nickName : [{
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

               loginCode : [{
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
            nickName: 'nick_name',
            roleGameLevel: 'role_game_level',
            unBoundIngotQuantity: 'un_bound_ingot_quantity',
            boundIngotQuantity: 'bound_ingot_quantity',
            totalIngotQuantity: 'total_ingot_quantity',
            unBoundTalisman: 'un_bound_talisman',
            boundTalisman: 'bound_talisman',
            totalTalisman: 'total_talisman',
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
  const table = await getGameUserList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  BannedStatusOptions.value = await getDictFunc('BannedStatus')
  OnlineStatusOptions.value = await getDictFunc('OnlineStatus')

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
            deleteGameUserFunc(row)
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
      const res = await deleteGameUserByIds({ IDs })
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
const updateGameUserFunc = async(row) => {
    const res = await findGameUser({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteGameUserFunc = async (row) => {
    const res = await deleteGameUser({ ID: row.ID })
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
  const res = await findGameUser({ ID: row.ID })
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
