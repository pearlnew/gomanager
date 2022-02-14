import MockAdapter from "axios-mock-adapter"
import axios from "axios"

export default {
    bootstrap() {
        var mock = new MockAdapter(axios)
        mock.onPost('/api/login').reply(200, {
            code: 200,
            msg: '',
            data: {
                username: 'pearl',
                gender:1,
                userid:1,
                registtime: '2020-01-01',
                isAdmin: true,
                age: 18,
                nickName: '女王',
                Province: '上海',
                City:'上海',
                avatarUrl: 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIrv2ibfCvqpGu4PJuZ1hQ1ydkibO7muTBUwoM9TcTBzs8CqHicCxZcLed9nTZ1XABhTtXGfk1o1lQyQ/132'
            }
        })

        mock.onGet('/api/user/list').reply(200, {
            code: 200,
            msg: '',
            data: {
                list: [
                    {username: 'pearl',gender:2, userid:1,registtime: '2020-01-01', isAdmin: true, age: 18, nickName: '女王', Province: '上海', City:'上海', avatarUrl: 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIrv2ibfCvqpGu4PJuZ1hQ1ydkibO7muTBUwoM9TcTBzs8CqHicCxZcLed9nTZ1XABhTtXGfk1o1lQyQ/132'},
                    {username: 'zhang3', gender:1, userid:2, registtime: '2022-01-01', isAdmin: false, age: 18, nickName: '张三', Province: '北京', City:'北京', avatarUrl: 'https://thirdwx.qlogo.cn/mmopen/vi_32/VkyFrnwibyhVPh1ejCXDZ9qHebsWBTlEn9wTuvibQGzLPEEWtIdu63uDNrBgsVQVr5RMicGuH16VwYGYWSVb2icJZg/132'},
                    {username: 'li4', gender:1, userid:3, registtime: '2021-01-01', isAdmin: false, age: 18, nickName: '里斯', Province: '广东', City:'珠海', avatarUrl: 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIrv2ibfCvqpGu4PJuZ1hQ1ydkibO7muTBUwoM9TcTBzs8CqHicCxZcLed9nTZ1XABhTtXGfk1o1lQyQ/132'},
                    {username: 'wang5', gender:1, userid:4, registtime: '2020-08-01', isAdmin: false, age: 18, nickName: '王五', Province: '天津', City:'天津', avatarUrl: 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJDdmHm6txEVuKcwEiaqHnQqR9iaKW2axhYkhHpbrH7K4cciaBTAyxWiaS2qvuSO52d0Dv3hwFPMWiaBlw/132'},
                ],
                pageTotal: 1
            }
        })

        mock.onGet('/api/user/1').reply(200, {
            code: 200,
            msg: '',
            data: {
                userInfo: {
                    username: 'pearl',
                    gender:1,
                    userid:1,
                    registtime: '2020-01-01',
                    isAdmin: true,
                    age: 18,
                    nickName: '女王',
                    Province: '上海',
                    City:'上海',
                    avatarUrl: 'https://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIrv2ibfCvqpGu4PJuZ1hQ1ydkibO7muTBUwoM9TcTBzs8CqHicCxZcLed9nTZ1XABhTtXGfk1o1lQyQ/132'
                },
                userPerms: []
            }
        })
    }
}