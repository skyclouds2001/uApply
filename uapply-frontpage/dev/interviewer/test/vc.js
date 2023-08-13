const test_user = {
	avatarUrl: "https://thirdwx.qlogo.cn/mmopen/vi_32/TEgslcia9JkwLhBdttuLA4DaBqTDIe8DQqXR7s8FsKAhXAxhRbDz0iaGkktn3OwuictKnnW3pJvkPClVT1eUzlEWg/132",
	city: "",
	country: "",
	gender: 0,
	language: "zh_CN",
	nickName: "cmtheit818",
	province: "",
	"accessToken":"aliqua adipisicing",
	"accessExpire":43150708.12135562,
	"refreshAfter":10000000000,
	"uid":83242572.53640354
}

const test_orgs = [
	{
		'id': 'alpha',
		'name': 'alpha',
		'department': [{
			dep_name: 'a1',
			dep_id: 'a1'
		}, {
			dep_name: 'a2',
			dep_id: 'a2'
		}, {
			dep_name: 'a3',
			dep_id: 'a3'
		}],
	}, {
		'id': 'beta',
		'name': 'beta',
		'department': [{
			dep_name: 'b1',
			dep_id: 'b1'
		}, {
			dep_name: 'b2',
			dep_id: 'b2'
		}, {
			dep_name: 'b3',
			dep_id: 'b3'
		} ],
	},
];
const test_resume = {
	"name": " hello",
	"sex": "男",
	"student_num": "dolor veniam in culpa sed",
	"address_num": "non ipsum",
	"major": "aliqua pariatur enim veniam",
	"phone_num": "ad incididunt nostru",
	"email": "do amet qui sit",
	"intro": "amet aute nisi commodo voluptateamet aute nisi commodo voluptateamet aute nisi commodo voluptateamet aute nisi commodo voluptateamet aute nisi commodo voluptate"
}

const NEW = 0,
ARRANGED = 1,
INTERVIEWED = 2,
PASS = 3,
OUT = 4,
test_dps =  [{
	dep_id: 'a1',
	dep_name: 'a1',
	org_id: 'alpha',
	org_name: 'alpha',
	first: NEW,
	second: PASS,
	final: NEW
}, {
	dep_id: 'a2',
	dep_name: 'a2',
	org_id: 'alpha',
	org_name: 'alpha',
	first: NEW,
	second: ARRANGED,
	final: NEW
}, {
	dep_id: 'b2',
	dep_name: 'b2',
	org_id: 'beta',
	org_name: 'beta',
	first: ARRANGED,
	second: NEW,
	final: NEW
}, {
	dep_id: 'b3',
	dep_name: 'b3',
	org_id: 'beta',
	org_name: 'beta',
	first: PASS,
	second: PASS,
	final: OUT
}]

/**
 * #  /app.js : get_user 函数第一行插入; 
 * #  虚拟用户信息
*/
export function vc(testUser=test_user){
	wx.setStorageSync('user', testUser)
	wx.setStorageSync('lst', Date.now())
}

export function uvc(){
	wx.setStorageSync('user', {})
	wx.setStorageSync('lst', 0)
}

export function vo(){
	return test_orgs
}

export function vdps(){
	return test_dps
}

export function vrsm(){
	return test_resume
}

export function uvrsm(){
	for(const i in test_resume){
		test_resume[i] = ""
	}
	return test_resume
}

export function toResume(){
    wx.navigateTo({
        url: "/subpages/resumes/resumes"
    })
}