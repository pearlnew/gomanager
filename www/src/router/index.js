import {createRouter, createWebHistory} from "vue-router";

export const routes = [
    {
        path: '/',
        redirect: '/dashboard.html'
    },{
        path: '/login.html',
        name: 'login',
        meta: {
            title: '登陆'
        },
        component: () => import('../views/Login.vue')
    },{
        path: '/',
        name: 'Home',
        component: () => import('../views/Home.vue'),
        children: [
            {
                path: '/dashboard.html',
                name: 'dashboard',
                meta: {
                    title: '系统首页',
                    isMenu: true,
                    icon: "el-icon-home-filled"
                },
                component: () => import('../views/Dashboard.vue')
            }
        ]
    }, {
        path: '/user',
        name: 'User',
        component: () => import('../views/Home.vue'),
        meta: {
            title: '用户管理',
            isMenu: true,
            icon: "el-icon-user-filled"
        },
        children: [
            {
                path: '/user/:id.html',
                name: 'userDetail',
                meta: {
                    perm: 'isAdmin',
                    title: '用户详情'
                },
                component: () => import('../views/UserDetail.vue')
            }, {
                path: '/user/list.html',
                name: 'userList',
                meta: {
                    perm: 'isAdmin',
                    isMenu: true,
                    title: '用户列表'
                },
                component: () => import('../views/UserList.vue')
            }
        ]
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

router.beforeEach((to, from, next) => {
    document.title = `${to.meta.title} | vue-manage-system`;
    const role = localStorage.getItem('userid');
    if ((!role || role === "0") && to.path !== '/login.html') {
        next('/login.html');
    } else if (to.meta.permission) {
        // 如果是管理员权限则可进入，这里只是简单的模拟管理员权限而已
        role === 'admin'
            ? next()
            : next('/403');
    } else {
        next();
    }
});

export default router;