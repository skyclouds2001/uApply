import http from '../http'

// 组织登录
export function OrganizationLogin(params) {
    return http.post("/org/login", params);
}

// 组织下部门登录
export function DepartmentLogin(params) {
    return http.post("/dep/login", params);
}

// 获取部门信息
export function GetAllDepartmentInfo() {
    return http.get("/org/deps");
}

// 设置组织招新日期
export function SetRecruitmentDate(params) {
    return http.post("/org/add-time", params);
}

// 获取组织招新日期
export function GetRecruitmentDate(proxy) {
    const { id } = proxy.$localData("get", "userInfo");
    return http.get("/org/enroll-time/" + id);
}

// 注册新部门
export function DepartmentRegistration(params) {
    return http.post("/org/cre-dep", params);
}

// 删除部门
export function DeleteDepartment(id) {
    return http.delete("/org/dep/" + id);
}

// 修改部门信息
export function EditDepartment(id, params) {
    return http.patch("/org/dep/" + id, params);
}

// 获取组织招新数据
export function GetOrganizationRecruitmentDate() {
    return http.get("/org/get-data");
}




// 配置面试官信息
export function AddInterviewer(params) {
    return http.post("/dep/add-inter", params);
}

// 获取所有面试官
export function GetAllInterviewers() {
    return http.get("/dep/inters");
}

// 删除面试官
export function DeleteInterviewer(ID) {
    return http.get("/dep/delete-Inter/" + ID);
}

// 修改面试官信息
export function EditInterviewer(params) {
    return http.post("/dep/updateInter", params);
}

// 获取部门投递数据
export function GetAllInterviewees() {
    return http.post("/dep/getusers-first");
}

// 获取一面名单
export function GetAllInterviewedOnce() {
    return http.post("/dep/getusers-first/interviewed");
}

// 获取淘汰名单
export function GetAllEliminate() {
    return http.post("/dep/getout");
}

// 获取部门招新数据
export function GetDepartmentData() {
    return http.get("/dep/reg-data");
}

// 查看录取名单
export function GetAdmission() {
    return http.post("/dep/getenroll");
}

// 淘汰人员
export function EliminateInterviewee(turn,params) {
    return http.post("/dep/out/" + turn, params);
}

// 批量通过面试
export function PassInterview(turn, params) {
    return http.post("/dep/pass/" + turn, params);
}

// 设置面试信息
export function SetRecruitmentConfig(turn, params) {
    return http.post("/dep/set-conf/" + turn, params);
}

// 手动删除已录取成员
export function DeleteAdmission(params) {
    return http.delete("/dep/del-enrolled", params);
}

// 手动添加录取成员
export function AddAdmission(params) {
    return http.post("/dep/add-enroll", params);
}

// 获取面试配置
export function GetRecruitmentConfig(turn) {
    return http.get("/dep/get-conf/" + turn);
}

// 录取
export function PassAndAdmission(params) {
    return http.post('/dep/enroll', params);
}

// 发送面试通知
export function SendNotification(turn, params) {
    return http.post('/dep/send-inter-msg/' + turn, params);
}