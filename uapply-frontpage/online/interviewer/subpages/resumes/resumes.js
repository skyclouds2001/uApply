import {uvrsm} from "../../test/vc";

const app = getApp()
Page({
	data: {
		app,
		leave: false,
		success_eva: false,
        num: 0,
        eva: 0,
        mark: ""
	},
	// 加载用户简历
	onLoad(){
		const thisPage = this;
		this.getOpenerEventChannel().on('id', (id) => {
			this.setData({
				...id // uid and depId
			})
			wx.showLoading({
				title: "加载用户简历中",
			})
			// 请求用户轮次
			app.request({
				url: {
					key: 'inter_turn'
				},
				accessToken: true,
				data: {
					uid: Number(this.data.uid),
					depId: Number(this.data.depId)
				},
				fail_back: true,
                // 请求用户简历
				success(res){
                    thisPage.setData({
                        turn: res.data.turn
                    })
					app.request({
						url: {
							key: 'inter_getresume',
							arg: String(res.data.turn)
                        },
                        accessToken: true,
						method: "POST",
						data: {
							uid: Number(thisPage.data.uid),
							depId: Number(thisPage.data.depId)
						},
						fail_back: true,
						success(res){
							wx.hideLoading()
							thisPage.setData({
								...res.data,
                            })
                            if(res.data.tag){
                                wx.showLoading({
                                  title: '获取面评',
                                }).then(() => {
                                    app.request({
                                        url: {
                                            key: 'inter_vieweva',
                                            arg: String(thisPage.data.num)
                                        },
                                        accessToken: true,
                                        data: {
                                            depId: Number(thisPage.data.depId),
                                            uid: Number(thisPage.data.uid),
                                        },
                                        fail_back: true,
                                        success(res){
                                            thisPage.setData({
                                                ...res.data
                                            })
                                        }
                                    })
                                })
                            }
						},
					})
				},
			})
		})
	},
	leave(){
		this.setData({
			leave: true
		})
	},
	back(){
		this.setData({
			leave: false
		})
	},
	submit(e){
		const thisPage = this;
		if (!(e.detail.value.eva && e.detail.value.mark)){
			wx.showToast({
				title: "请完成面试评价并打分",
				icon: "none",
				duration: 3000
			})
		} else{
            const thispage = this;
            wx.showLoading({
                title: '提交中',
                success(){
                    app.request({
                        url:{
                            key: 'inter_evaluate',
                            arg: String(thisPage.data.turn)
                        },
                        accessToken: true,
                        data: {
                            ...e.detail,
                            uid: Number(thisPage.data.uid),
                            depId: Number(thispage.data.depId),
                        },
                        success(res){
                            wx.hideLoading({
                                success(){
                                    thisPage.setData({
                                        ...res.data,
                                        success_eva: true
                                    })
                                }
                            })
                        }
                    })
                }
            })
			
		}
	}
})