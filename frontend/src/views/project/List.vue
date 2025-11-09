<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {Project, ProjectStatus} from '@/utils/tables';
import {useUrlSearchParams} from '@vueuse/core';
import {reactive, ref, watchEffect} from 'vue';
import {Plus, Search} from '@element-plus/icons-vue'
import {useRouteParams} from '@vueuse/router';

const projects = ref<{
  data: Project[],
  total: number
} | null>(null)


const id = useRouteParams('id')
const params = useUrlSearchParams<Record<string, string>>('history')
watchEffect(async () => projects.value = await request('GET', id.value ? `/clients/${id.value}/projects` : '/projects', params))

const projectStatuses = ref<ProjectStatus[]>([])
request<ProjectStatus[]>('GET', '/project-statuses').then(res => projectStatuses.value = res)

</script>

<template>

  <div class="p-4 flex flex-col gap-4">

    <div class="text-xl">{{$t('clientDetail.projectsTitle')}}</div>

    <div class="flex flex-wrap items-center gap-4">
      <el-input
        v-model="params.name"
        :prefix-icon="Search"
        :placeholder="$t('projectList.searchPlaceholder')"
        class="!w-64"
      />
      <el-button @click="delete params.statusId">{{$t('projectList.statusAll')}}</el-button>
      <el-segmented v-model="params.statusId" :options="projectStatuses" :props="{ value: 'id' }">
        <template #default="scope">
          <div class="flex flex-col items-center">
            <div>{{scope.item.projectCount}}</div>
            <div>{{scope.item.name}}</div>
          </div>
        </template>
      </el-segmented>
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

  </div>

</template>
