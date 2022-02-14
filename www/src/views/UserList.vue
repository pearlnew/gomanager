<template>
  <div>
    <div class="container">
      <div class="handle-box">
        <el-select v-model="query.address" placeholder="地址" class="handle-select mr10">
          <el-option key="1" label="广东省" value="广东省"></el-option>
          <el-option key="2" label="湖南省" value="湖南省"></el-option>
        </el-select>
        <el-input v-model="query.username" placeholder="用户名" class="handle-input mr10"></el-input>
        <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
      </div>
      <el-table
          :data="tableData"
          border
          class="table"
          ref="multipleTable"
          header-cell-class-name="table-header"
      >
        <el-table-column prop="UserID" label="ID" width="55" align="center">
          <template #default="scope">
            <a target="_self" :href="goToUserDetail(scope.row.userid)">{{ scope.row.userid }}</a>
          </template>
        </el-table-column>
        <el-table-column label="头像(查看大图)" align="center">
          <template #default="scope">
            <el-image
                class="table-td-thumb"
                :src="scope.row.avatarUrl"
            ></el-image>
          </template>
        </el-table-column>
        <el-table-column prop="nickName" label="用户名"></el-table-column>
        <el-table-column label="性别">
          <template #default="scope"><el-tag
              :type="scope.row.gender === 2 ? 'success' : 'danger'"
          >{{ scope.row.gender === 2 ? '女': '男' }}</el-tag></template>
        </el-table-column>
        <el-table-column label="地址">
          <template #default="scope">{{scope.row.Province}} {{scope.row.City}}</template>
        </el-table-column>
        <el-table-column label="管理员" align="center">
          <template #default="scope"><el-tag
              :type="scope.row.isAdmin === 1 ? 'success' : ''"
          >{{ scope.row.isAdmin === 1 ? '是': '否' }}</el-tag></template>
        </el-table-column>

        <el-table-column label="注册时间">
          <template #default="scope">
            <span>{{ parseTime(scope.row.registtime, '{y}-{m}-{d}') }}</span>
          </template></el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
            background
            layout="total, prev, pager, next"
            :current-page="query.pageIndex"
            :page-size="query.pageSize"
            :total="pageTotal"
            @current-change="handlePageChange"
        ></el-pagination>
      </div>
    </div>

    <!-- 编辑弹出框 -->
    <el-dialog title="编辑" v-model="editVisible" width="30%">
      <el-form ref="form" :model="form" label-width="70px">
        <el-form-item label="用户名">
          <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="form.address"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
                <span class="dialog-footer">
                    <el-button @click="editVisible = false">取 消</el-button>
                    <el-button type="primary" @click="saveEdit">确 定</el-button>
                </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { getUserList } from "../api/index";
import {parseTime} from "../utils/index"
export default {
  username: "basetable",
  data() {
    return {
      query: {
        address: "",
        username: "",
        pageIndex: 1,
        pageSize: 10
      },
      tableData: [],
      multipleSelection: [],
      delList: [],
      editVisible: false,
      pageTotal: 0,
      form: {},
      idx: -1,
      id: -1
    };
  },
  created() {
    this.getData();
  },
  methods: {
    parseTime,
    getData() {
      getUserList(this.query).then(res => {
        console.log(res);
        if (res.data.code != 200 ){
          console.log(res.data.msg)
        } else {
          this.tableData = res.data.data.list;
          this.pageTotal = res.data.data.pageTotal || 1;
        }
      });
    },
    // 触发搜索按钮
    handleSearch() {
      this.$set(this.query, "pageIndex", 1);
      this.getData();
    },
    // 分页导航
    handlePageChange(val) {
      this.$set(this.query, "pageIndex", val);
      this.getData();
    },

    goToUserDetail(userid) {
      return "/user/"+userid+".html"
    }
  }
};
</script>

<style scoped>
.handle-box {
  margin-bottom: 20px;
}

.handle-select {
  width: 120px;
}

.handle-input {
  width: 300px;
  display: inline-block;
}
.table {
  width: 100%;
  font-size: 14px;
}
.red {
  color: #ff0000;
}
.mr10 {
  margin-right: 10px;
}
.table-td-thumb {
  display: block;
  margin: auto;
  width: 40px;
  height: 40px;
}
</style>
