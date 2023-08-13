// pages/registors.js

const test_resume = {

}
const app = getApp()
Page({
	data: {
		app,
        have_emitted: {
            total: null,
            male: null,
            female: null
        },
        have_passed: {
            total: null,
            male: null,
            female: null
        }
	},
	onLoad(){
        const thisPage = this;
        this.getOpenerEventChannel().on('getData', (d) => {
            this.setData({
                ...d
            })
            app.request({
                url: {
                    key: 'inter_get-data',
                    arg: String(d.depId)
                },
                accessToken: true,
                fail_back: true,
                success(res){
                    wx.hideLoading({
                        success(){
                            thisPage.setData({
                                have_emitted: {
                                    total: res.data.sum,
                                    male: res.data.sum_male,
                                    female: res.data.sum_female
                                },
                                have_passed: {
                                    total: res.data.pass,
                                    male: res.data.pass_male,
                                    female: res.data.pass_female
                                }
                            })
                        }
                    })
                }
            })
        })
        wx.showLoading({
            title: '加载中',
        })

	},
})