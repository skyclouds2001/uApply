import http from "../http"

// 用户通过学校统一登录接口登录
export function UserLogin(data) {
    return http.post("/user/login-xd", data);
}

// 查看简历是否存在/查看简历
export function GetResume() {
    return http.get('/user/check-resume');
}

// 用户填写简历
export function SaveResume(data){
    return http.post('/user/resume',data);
}

// 获取招新的组织
export function GetOrganization(){
    return http.get('/user/org-enrolling');
}

// 获取招新的部门
export function GetDepartment(id){
    return http.get('/user/org-deps/' + id);
}

// 用户报名
export function SubmitRegistration(data){
    return http.post('/user/register',data);
}

// 获取报名状态
export function GetStatus(){
    return http.get('/user/register');
}