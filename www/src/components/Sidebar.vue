<template>
  <div class="sidebar">
    <el-menu class="sidebar-el-menu" :default-active="onRoutes"  background-color="#324157"
             text-color="#bfcbd9" active-text-color="#20a0ff" unique-opened router>
      <template v-for="item in items" :key="item.path">
        <div v-if="item.meta.perm === undefined || hasPerm(item.meta.perm)">
          <template v-if="item.children">
            <el-sub-menu :index="item.path" :key="item.path" >
              <template #title>
                <el-icon color="white"><component :is="item.meta.icon"></component></el-icon>
                <span >{{ item.meta.title }}</span>
              </template>
              <template v-for="subItem in item.children" :key="subItem.path">
                <div v-if="subItem.meta.perm === undefined || hasPerm(subItem.meta.perm)">
                  <el-menu-item :index="subItem.path">{{ subItem.meta.title }}</el-menu-item>
                </div>
              </template>
            </el-sub-menu>
          </template>
          <template v-else>
            <el-menu-item :index="item.path" :key="item.path">
              <el-icon color="white"><component :is="item.meta.icon"></component></el-icon>
              <template #title>{{ item.meta.title }}</template>
            </el-menu-item>
          </template>
        </div>
      </template>
    </el-menu>
  </div>
</template>

<script>
import {routes} from '../router'

export default {
  data() {
    return {
      items: [],
      leftmenu: routes
    }
  },
  mounted() {
    var menus = []
    for (let route of this.leftmenu) {
      var item = {}
      const subItem = []
      if (route.meta && route.meta.isMenu) {
        item = route
      }
      if (route.children){
        for (let subRoute of route.children) {
          if (subRoute.meta.isMenu) {
            subItem.push(subRoute)
          }
        }
      }
      if (this.isEmpty(item)) {
        if (subItem.length > 0) {
          menus.push(subItem[0])
        }
      } else if (subItem.length > 0){
        const t = item
        t.children = subItem
        menus.push(t)
      }
    }
    console.log(menus)
    this.items = menus
  },
  computed: {
    onRoutes() {
      return this.$route.path.replace("/", "");
    }
  },
  methods: {
    hasPerm(name) {
      return true
      // if(name === 'isAdmin') {
      //   const user = JSON.parse(localStorage.getItem("user"))
      //   return (user.isAdmin === 1)
      // } else {
      //   const perms = JSON.parse(localStorage.getItem("perms"))
      //   for (let perm of perms) {
      //     return perm.PermCode === name
      //   }
      // }
    },
    isEmpty(ob) {
      for(let i in ob){ return false;}
      return true;
    }
  }
};
</script>

<style>
.sidebar {
  display: block;
  position: absolute;
  left: 0;
  top: 60px;
  bottom: 0;
  overflow-y: scroll;
}
.sidebar::-webkit-scrollbar {
  width: 0;
}
.sidebar-el-menu:not(.el-menu--collapse) {
  width: 150px;
}
.sidebar > ul {
  height: 100%;
}

.sidebar .el-menu-item {
  padding-left: 0 !important;
  font-size: 10px;
}
.sidebar .el-sub-menu__title {
  padding-left: 0 !important;
  font-size: 10px;
}

.sidebar .el-menu--inline .el-menu-item {
  padding-left: 40px !important;
  font-size: 10px;
  color: rgb(191, 203, 217) !important;
}

</style>
