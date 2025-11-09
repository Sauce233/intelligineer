<script setup lang="ts">
import {useSessionStore} from '@/stores/session';
import {request} from '@/utils/axios';
import {useRouteParams} from '@vueuse/router';
import {reactive} from 'vue';


export interface ProjectForm {
  name: string;
  offer: number;
  budget: number;
  statusId: string;
  profile: string;
}

const id = useRouteParams('id')
const session = useSessionStore()
const projectForm = reactive({
  name: '', profile: '', offer: 0, budget: 0, statusId: '',
})
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <el-form :model="projectForm">
      <el-form-item :label="$t('name')">
        <el-input v-model="projectForm.name">
        </el-input>
      </el-form-item>
      <el-form-item :label="$t('offer')">
        <el-input-number v-model="projectForm.offer">
          <template #prefix>¥</template>
        </el-input-number>
      </el-form-item>
      <el-form-item :label="$t('budget')">
        <el-input-number v-model="projectForm.budget">
          <template #prefix>¥</template>
        </el-input-number>
      </el-form-item>
      <el-form-item :label="$t('status')">
        <el-select v-model="projectForm.statusId" class="!w-64">
          <el-option v-for="t in session.enums['project_statuses']" :key="t.id" :label="t.name" :value="t.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('profile')">
        <el-input v-model="projectForm.profile" type="textarea" autosize>
        </el-input>
      </el-form-item>
    </el-form>
      <el-button type="primary"
        @click="request('POST', `/clients/${id}/projects`, projectForm).then(id => id && $router.push(`/projects/${id}`))"
      >
        {{$t('confirm')}}
      </el-button>
  </div>
</template>
