<script setup lang="ts">
import {request} from '@/utils/axios';
import {Setting} from '@element-plus/icons-vue';
import {reactive, ref, watch} from 'vue';
import {useI18n} from 'vue-i18n';
import {useRoute, useRouter} from 'vue-router';

const { t } = useI18n({ messages: {
  zh: {
    login: '登录',
    retrieve: '找回密码',
    usernameOrEmail: '用户名 / 邮箱',
    email: '邮箱',
    password: '密码',
    authcode: '验证码',
    oldPassword: '旧密码',
    newPassword: '新密码',
    resetPassword: '更改密码',
  },
} })

const router = useRouter()
const route = useRoute()
const currentTab = ref(route.path)
watch(currentTab, tab => router.push(tab))

const loginForm = reactive({
  email: '',
  password: '',
})

async function login() {
  if (await request('POST', '/login', undefined, loginForm)) {
    router.push('/dashboard')
  }
}

const retrieveForm = reactive({
  email: '',
  authcode: '',
  newPassword: '',
})
</script>

<template>
  <div class="h-full flex justify-center items-center">
    <div class="card flex flex-col gap-4">
      <div class="flex justify-between mx-4">
        <div class="text-xl">
          {{t('login')}}
        </div>
        <el-button :icon="Setting" @click="$router.push('/setting')" circle>
        </el-button>
      </div>
      <el-tabs v-model="currentTab">

        <el-tab-pane :label="t('login')" name="/login">
          <el-form :model="loginForm" label-width="auto">
            <el-form-item :label="t('email')">
              <el-input v-model="loginForm.email" class="w-64">
              </el-input>
            </el-form-item>
            <el-form-item :label="t('password')">
              <el-input
                v-model="loginForm.password"
                class="w-64"
                type="password"
                @keyup.enter="login"
              ></el-input>
            </el-form-item>
            <el-form-item>
              <el-button class="mx-auto" @click="login">
                {{t('confirm')}}
              </el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>

        <el-tab-pane :label="t('retrieve')" name="/retrieve">
          <el-form :model="retrieveForm" label-width="auto">
            <el-form-item :label="t('email')">
              <el-input v-model="retrieveForm.email" class="w-64">
              </el-input>
            </el-form-item>
            <el-form-item :label="t('authcode')">
              <el-input v-model="retrieveForm.authcode" class="w-64">
              </el-input>
            </el-form-item>
            <el-form-item :label="t('newPassword')">
              <el-input v-model="retrieveForm.newPassword" class="w-64">
              </el-input>
            </el-form-item>
            <el-form-item>
              <el-button class="mx-auto">
                {{t('confirm')}}
              </el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>

      </el-tabs>
    </div>
  </div>
</template>
