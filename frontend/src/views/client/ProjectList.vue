<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {Project, ProjectStatus} from '@/utils/tables';
import {useUrlSearchParams} from '@vueuse/core';
import {reactive, ref, watchEffect} from 'vue';
import {Plus, Search} from '@element-plus/icons-vue'
import {useRouteParams} from '@vueuse/router';
import type {List} from '@/utils/utils';
import ProjectEditor, {type Form} from '@/components/editor/ProjectEditor.vue';

const projects = ref<List<Project> | null>(null)

const id = useRouteParams('id')
const params = useUrlSearchParams<any>('history')
watchEffect(async () => projects.value = await request('GET', `/clients/${id.value}/projects`, params))

const projectStatuses = ref<ProjectStatus[]>([])
request<ProjectStatus[]>('GET', '/project-statuses').then(res => projectStatuses.value = res)
const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', profile: '', offer: 0, budget: 0, statusId: 0,
})

</script>

<template>

  <div class="text-xl flex justify-between">
    <div>{{$t('clientDetail.projectsTitle')}}</div>
    <el-button :icon="Plus" @click="isDialogOpen = true">
      {{$t('clientDetail.addProject')}}
    </el-button>
  </div>

  <div class="flex gap-4">
    <el-input v-model="params.name" :prefix-icon="Search" :placeholder="$t('projectList.searchPlaceholder')">
    </el-input>
  </div>
  <div v-if="projects" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="p in projects.data" :key="p.id" @click="$router.push(`/projects/${p.id}`)" class="card cursor-pointer flex flex-col gap-2">
      <div class="space-x-2 line-clamp-1">
        <tag :color="p.status.color">
          {{p.status.name}}
        </tag>
        <span class="text-lg">{{p.name}}</span>
      </div>
      <div class="text-subtle">{{p.client.name}}</div>
      <div class="flex flex-wrap gap-1">
        <tag v-for="t in p.tasks" :key="t.id" :color="t.status.color" class="!text-xs">
          {{t.name}}
        </tag>
      </div>
      <div class="mt-auto flex gap-2">
        <div class="text-subtle">{{$t('projectList.attachmentCount', { v: p.attachmentCount })}}</div>
        <div class="text-subtle">{{$t('projectList.invoiceCount', { v: p.invoiceCount })}}</div>
        <div class="text-subtle">{{$t('projectList.taskCount', { v: p.taskCount })}}</div>
      </div>
    </div>
  </div>
  <el-pagination v-if="projects" class="card" layout="prev, pager, next, total" :total="projects.total"
    :current-page="Number(params.page || 1)" :page-size="Number(params.size || 9)"
    @update:current-page="v => params.page = String(v)" @update:page-size="v => params.size = String(v)"
  />


  <el-dialog v-model="isDialogOpen" :title="$t('client.project.dialog.title')">
    <project-editor v-model="form">
    </project-editor>
    <template #footer>
      <el-button
        type="primary"
        @click="request('POST', `/clients/${id}/projects`, form).then(id => id && $router.push(`/projects/${id}`))"
      >
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>
</template>
