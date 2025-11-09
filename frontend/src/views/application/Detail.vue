<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import { useSessionStore } from '@/stores/session';
import { request } from '@/utils/axios';
import type { Application } from '@/utils/tables';
import { formatDate, getColorClass } from '@/utils/utils';
import { Back, Edit, Delete, Check, Close } from '@element-plus/icons-vue';
import { ref } from 'vue';
import { useRoute} from 'vue-router';

const application = ref<Application | null>(null)
const route = useRoute()
const session = useSessionStore()

request<Application>('GET', `/applications/${route.params.id}`).then(res => application.value = res)
</script>

<template>
  <div v-if="application" class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <!-- 顶部导航栏 -->
    <div class="card text-lg flex items-center flex-wrap gap-2">
      <el-button :icon="Back" @click="$router.back">{{$t('back')}}</el-button>
      <span>{{$t('application')}}: {{application.title}}</span>
      <tag :color="application.status.color">{{application.status.name}}</tag>
      <tag :color="application.type.color">{{application.type.name}}</tag>
      <div class="ms-auto flex gap-2">
        <el-button :icon="Edit" v-if="application.applicant.id === session.user?.id">
          {{$t('edit', { name: $t('application') })}}
        </el-button>
        <el-button :icon="Delete" v-if="application.applicant.id === session.user?.id">
          {{$t('delete', { name: $t('application') })}}
        </el-button>
      </div>
    </div>

    <!-- 基本信息 -->
    <div class="flex gap-4 flex-wrap">
      <div v-for="v, k in {
           applicant: application.applicant.name,
           approver: application.approver?.name || $t('pending'),
           createdAt: formatDate(application.createdAt),
           updatedAt: formatDate(application.updatedAt) }" :key="k"
        class="flex gap-2 items-center"
      >
        <div class="text-sm-subtle">{{$t(k)}}</div>
        <div>{{v}}</div>
      </div>
    </div>

    <!-- 申请内容 -->
    <div class="space-y-2">
      <div class="text-xl font-bold">{{$t('content')}}</div>
      <div class="bg-subtle p-4 rounded-2xl whitespace-pre-wrap">{{application.content}}</div>
    </div>

    <!-- 审批回复 -->
    <div v-if="application.reply" class="space-y-2">
      <div class="text-xl font-bold">{{$t('reply')}}</div>
      <div :class="['p-4 rounded-2xl whitespace-pre-wrap', getColorClass('bg', application.status.color)]">
        {{application.reply}}
      </div>
    </div>

    <!-- 审批操作 -->
    <div v-if="application.status.id === 1 && session.user?.id === application.approver?.id"
         class="card flex gap-4 bg-blue-50 dark:bg-blue-950">
      <div class="text-lg font-medium">{{$t('approvalAction')}}</div>
      <el-button type="success" :icon="Check" class="ms-auto">
        {{$t('approve')}}
      </el-button>
      <el-button type="danger" :icon="Close">
        {{$t('reject')}}
      </el-button>
    </div>
  </div>
</template>
