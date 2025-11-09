<script setup lang="ts">
import {useSessionStore} from '@/stores/session';
import {api} from '@/utils/axios';
import {formatDate} from '@/utils/format-date';
import {Edit, EditPen, Key, User} from '@element-plus/icons-vue';
import {reactive, ref} from 'vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  'zh-CN': {
    userInfo: '用户信息',
    reLogin: '重新登录',
    myProjects: '我加入的项目',
    myApplies: '我发出的申请',
    writeApply: '写申请',
  },
} })

const session = useSessionStore()
async function editUserInfo() {

}

const isApplyDialogOpen = ref(false)
const applyForm = reactive({
  title: '',
  content: '',
  userApplyTypeId: '',
})
async function addApply() {
  await api.post('/user-applies', applyForm)
  isApplyDialogOpen.value = false
}

</script>

<template>
  <div v-if="session.user" class="max-w-5xl mx-auto space-y-4 p-4">

    <div v-if="session.user" class="card space-y-4">
      <div class="mx-4 text-lg">
        {{t('userInfo')}}
      </div>
      <div class="flex flex-col gap-2">
        <lite-desc :label="t('registeredIn')">
          {{formatDate(session.user.createdAt)}}
        </lite-desc>
        <lite-desc :label="t('updatedAt')">
          {{formatDate(session.user.updatedAt)}}
        </lite-desc>
        <lite-desc :label="t('email')">
          {{session.user.email}}
        </lite-desc>
      </div>
      <name-card :user="session.user">
      </name-card>
      <div class="flex gap-2">
        <el-button @click="$router.push('/login')" :icon="User">
          {{t('reLogin')}}
        </el-button>
        <el-button @click="editUserInfo" :icon="Edit">
          {{t('edit')}}
        </el-button>
        <el-button @click="$router.push('/retrieve')" :icon="Key">
          {{t('retrieve')}}
        </el-button>
      </div>
    </div>

    <div class="card space-y-4">
      <div class="mx-4 text-lg">
        {{t('myProjects')}}
      </div>
      <div>
<!--
        <div
          v-for="up in session.user.userProjects"
          :key="up.projectId"
          class="table-item space-y-2"
          @click="$router.push(`/projects/${up.projectId}`)"
        >
          <div class="flex justify-between">
            <lite-desc :label="t('name')">
              {{up.project.name}}
            </lite-desc>
            <lite-desc :label="t('hours')">
              {{up.hours}}{{t('hour')}}
            </lite-desc>
          </div>
          <lite-desc :label="t('duty')" class="text-sm-subtle">
            {{up.duty}}
          </lite-desc>
        </div>
-->
      </div>
    </div>

    <div class="card space-y-4">
      <div class="mx-4 flex justify-between">
        <div class="text-xl">{{t('myApplies')}}</div>
        <el-button :icon="EditPen" @click="isApplyDialogOpen = true">{{t('writeApply')}}</el-button>
      </div>
      <el-dialog v-model="isApplyDialogOpen" :title="t('writeApply')" class="w-dialog">
        <el-form :model="applyForm" label-width="auto">
          <el-form-item :label="t('title')">
            <el-input v-model="applyForm.title" class="w-64">
            </el-input>
          </el-form-item>
          <el-form-item :label="t('content')">
            <el-input v-model="applyForm.content" type="textarea">
            </el-input>
          </el-form-item>
          <el-form-item :label="t('userApplyType')">
            <el-select v-model="applyForm.userApplyTypeId">
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="addApply">
            {{t('confirm')}}
          </el-button>
        </template>
      </el-dialog>
      <div>
        <!--
        <div
          v-for="ua in session.user.userApplies"
          :key="ua.id"
          class="table-item"
          @click="$router.push(`/user-applies/${ua.id}`)"
        >
          <div class="flex justify-between items-center">
            <div class="text-xl space-x-2">
              <span>
                {{ua.title}}
              </span>
              <el-tag :type="userApplyTypes[ua.type].type">
                {{userApplyTypes[ua.type].label}}
              </el-tag>
              <el-tag v-if="ua.status" type="success">
                {{t('solved')}}
              </el-tag>
              <el-tag v-else type="primary">
                {{t('pending')}}
              </el-tag>
            </div>
            <div class="text-sm">
              <lite-desc :label="t('createdAt')">{{formatDate(ua.createdAt)}}</lite-desc>
              <lite-desc :label="t('updatedAt')">{{formatDate(ua.updatedAt)}}</lite-desc>
            </div>
          </div>
          <div class="text-sm-subtle">
            <span class="font-bold">{{t('reply')}}: </span>{{ua.reply}}
          </div>
        </div>
        -->
      </div>
    </div>

  </div>
</template>
