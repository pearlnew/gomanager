<template>
  <el-row :gutter="20">
    <el-col :span="7">
      <el-card shadow="hover" class="mgb20">
        <div class="user-info">
          <img :src="userInfo.avatarUrl" class="user-avator" alt />
          <div class="user-info-cont">
            <div class="user-info-name">{{ userInfo.nickName }}</div>
            <div>{{userInfo.isAdmin === 1 ? '管理员': '普通用户'}}</div>
          </div>
        </div>
        <el-row class="user-info-list row-bg">
          <el-col :span="10" style="text-align: right">注册时间：</el-col>
          <span>{{ parseTime(userInfo.registtime, '{y}-{m}-{d}')}}</span>
        </el-row>
        <el-row class="user-info-list row-bg">
          <el-col :span="10" style="text-align: right">性别：</el-col>
          <span>{{ userInfo.gender === 2 ? '女':'男' }}</span>
        </el-row>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import {getUserDetail} from "../api/index"
import {parseTime} from "../utils/index"
export default {
  name: "UserDetail",
  data() {
    return {
      userInfo: {},
      userPerms: []
    }
  },
  mounted() {
    this.getUserInfo()
  },
  methods: {
    parseTime,
    getUserInfo() {
      let userid = this.$route.params.id
      getUserDetail(userid).then(res => {
        if(res.status === 200 && res.data.code===200) {
          this.userInfo = res.data.data.userInfo;
          this.userPerms = res.data.data.userPerms;
        } else {
          console.log("error")
        }
      })
    }
  }
}
</script>

<style scoped>
.user-info {
  display: flex;
  align-items: center;
  padding-bottom: 10px;
  border-bottom: 2px solid #ccc;
  margin-bottom: 10px;
}
.user-avator {
  width: 120px;
  height: 120px;
  border-radius: 50%;
}

.user-info-cont {
  padding-left: 15px;
  flex: 1;
  font-size: 14px;
  color: #999;
}
.user-info-cont div:first-child {
  font-size: 30px;
  color: #222;
}

.user-info-list {
  font-size: 14px;
  color: #999;
  line-height: 30px;
}
.red {
  color: #ff0000;
}
.mgb20 {
  margin-bottom: 20px;
}
</style>