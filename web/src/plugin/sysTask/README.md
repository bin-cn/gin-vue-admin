## 任务管理界面配置

### 1.代码存放目录
![代码目录.png](img%2F%E4%BB%A3%E7%A0%81%E7%9B%AE%E5%BD%95.png)
### 2. 配置

#### 2.1 字典配置
![定时任务状态.png](img%2F%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1%E7%8A%B6%E6%80%81.png)
![状态字典列表.png](img%2F%E7%8A%B6%E6%80%81%E5%AD%97%E5%85%B8%E5%88%97%E8%A1%A8.png)
#### 2.2 菜单配置
![菜单配置sysTask.png](img%2F%E8%8F%9C%E5%8D%95%E9%85%8D%E7%BD%AEsysTask.png)
![菜单配置sysTaskLog.png](img%2F%E8%8F%9C%E5%8D%95%E9%85%8D%E7%BD%AEsysTaskLog.png)

#### 2.3. 添加个函数，放入目录 utils/format.js
    export const truncateLength = (description) => {
        const maxLength = 10
        return description.length > maxLength
        ? description.slice(0, maxLength) + '...'
        : description
    }
#### 2.4 角色配置
![角色菜单管理.png](img%2F%E8%A7%92%E8%89%B2%E8%8F%9C%E5%8D%95%E7%AE%A1%E7%90%86.png)
![角色api配置.png](img%2F%E8%A7%92%E8%89%B2api%E9%85%8D%E7%BD%AE.png)

### 3.配置完成后启动系统前后端

### 4.添加任务，定时任务运行情况
![任务管理.png](img%2F%E4%BB%BB%E5%8A%A1%E7%AE%A1%E7%90%86.png)
![任务运行日志.png](img%2F%E4%BB%BB%E5%8A%A1%E8%BF%90%E8%A1%8C%E6%97%A5%E5%BF%97.png)
