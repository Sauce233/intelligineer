import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/login', component: () => import('@/views/Login.vue') },
    { path: '/settings', component: () => import('@/views/Setting.vue') },
    { path: '/', component: () => import('@/views/NavbarLayout.vue'), children: [
      { path: 'home', component: () => import('@/views/Home.vue') },
      { path: 'dashboard', component: () => import('@/views/Dashboard.vue') },
      { path: 'tasks', children: [
        { path: '', component: () => import('@/views/task/List.vue') },
        { path: ':id', component: () => import('@/views/task/Detail.vue'), children: [
          { path: '', component: () => import('@/views/task/MaterialList.vue') },
          { path: 'records', component: () => import('@/views/task/RecordList.vue') },
        ] },
      ] },
      { path: 'categories', children: [
        { path: '', component: () => import('@/views/category/List.vue') },
        { path: ':id', component: () => import('@/views/category/Detail.vue'), children: [
          { path: '', component: () => import('@/views/category/MaterialList.vue') },
        ] },
      ] },
      { path: 'applications', children: [
        { path: '', component: () => import('@/views/application/List.vue') },
        { path: ':id', component: () => import('@/views/application/Detail.vue') },
      ] },
      { path: 'purchases', children: [
        { path: '', component: () => import('@/views/purchase/List.vue') },
        { path: ':id', component: () => import('@/views/purchase/Detail.vue'), children: [
          { path: '', component: () => import('@/views/purchase/MaterialList.vue') },
        ] },
      ] },
      { path: 'records', children: [
        { path: '', component: () => import('@/views/record/List.vue') },
        { path: ':id', component: () => import('@/views/record/Detail.vue'), children: [
          { path: '', component: () => import('@/views/record/PhotoList.vue') },
        ] },
      ] },
      { path: 'clients', children: [
        { path: '', component: () => import('@/views/client/List.vue') },
        { path: ':id', component: () => import('@/views/client/Detail.vue'), children: [
          { path: '', component: () => import('@/views/client/ContactList.vue') },
          { path: 'projects', component: () => import('@/views/client/ProjectList.vue') },
        ] },
      ] },
      { path: 'projects', children: [
        { path: '', component: () => import('@/views/project/List.vue') },
        { path: ':id', component: () => import('@/views/project/Detail.vue'), children: [
          { path: '', component: () => import('@/views/project/AttachmentList.vue') },
          { path: 'tasks', component: () => import('@/views/project/TaskList.vue') },
          { path: 'invoices', component: () => import('@/views/project/InvoiceList.vue') },
        ] },
      ] },
      { path: 'users', children: [
        { path: '', component: () => import('@/views/user/List.vue') },
        { path: ':id', component: () => import('@/views/user/Detail.vue'), children: [
          { path: '', component: () => import('@/views/user/ProfileList.vue') },
          { path: 'records', component: () => import('@/views/user/RecordList.vue') },
          { path: 'applications', component: () => import('@/views/user/ApplicationList.vue') },
          { path: 'attendances', component: () => import('@/views/user/AttendanceList.vue') },
        ] },
      ] },
      { path: 'suppliers', children: [
        { path: '', component: () => import('@/views/supplier/List.vue') },
        { path: ':id', component: () => import('@/views/supplier/Detail.vue'), children: [
          { path: '', component: () => import('@/views/supplier/ContactList.vue') },
          { path: 'purchases', component: () => import('@/views/supplier/PurchaseList.vue') },
        ] },
      ] },
      { path: 'invoices', children: [
        { path: '', component: () => import('@/views/invoice/List.vue') },
        { path: ':id', component: () => import('@/views/invoice/Detail.vue') },
      ] },
      { path: 'materials', children: [
        { path: ':id', component: () => import('@/views/material/Detail.vue'), children: [
          { path: '', component: () => import('@/views/material/PurchaseList.vue') },
          { path: 'tasks', component: () => import('@/views/material/TaskList.vue') },
        ] },
      ] },
      { path: 'attendances', children: [
        { path: '', component: () => import('@/views/attendance/List.vue') },
        { path: ':id', component: () => import('@/views/attendance/Detail.vue'), children: [
          { path: '', component: () => import('@/views/attendance/UserList.vue') },
        ] },
      ] },
    ] },
  ],
})

export default router
