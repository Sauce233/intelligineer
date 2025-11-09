<script setup lang="ts">
import type {Form} from '@/components/editor/TaskEditor.vue';
import TaskEditor from '@/components/editor/TaskEditor.vue';
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Task } from '@/utils/tables';
import { formatDate, type List } from '@/utils/utils';
import { Plus, Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import { reactive, ref, watchEffect } from 'vue';

const tasks = ref<List<Task> | null>(null)

const id = useRouteParams('id')
const params = useUrlSearchParams<Record<string, string>>('history')
watchEffect(async () => tasks.value = await request('GET', `/projects/${id.value}/tasks`, params))

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', profile: '', time: [new Date(), new Date()], statusId: 0,
})

</script>

<template>

  <div class="flex justify-between">
    <div class="text-xl">
      {{$t('taskList.title')}}
    </div>
    <el-button :icon="Plus" @click="isDialogOpen = true">
      {{$t('taskList.addTaskButton')}}
    </el-button>
  </div>

  <div class="flex gap-4">
    <el-input v-model="params.name" class="!w-64" :prefix-icon="Search" :placeholder="$t('taskList.searchPlaceholder')">
    </el-input>
    <el-date-picker type="datetimerange"
      :model-value="params.timeRange"
      @update:model-value="v => params.timeRange = v?.map((e: Date) => e.toISOString())"
      :start-placeholder="$t('taskList.selectStartTime')"
      :end-placeholder="$t('taskList.selectEndTime')"
    ></el-date-picker>
  </div>

  <div v-if="tasks" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="t in tasks.data" :key="t.id" @click="$router.push(`/tasks/${t.id}`)" class="card cursor-pointer flex flex-col gap-2">
      <div class="space-x-2 line-clamp-1">
        <tag :color="t.status.color">
          {{t.status.name}}
        </tag>
        <span class="text-lg">{{t.name}}</span>
      </div>
      <div class="text-subtle link" @click.stop="$router.push(`/projects/${t.project.id}`)">{{t.project.name}}</div>
      <div class="text-subtle">{{t.profile}}</div>
      <div class="mt-auto flex flex-col gap-1">
          <div class="text-subtle">{{$t('taskList.materialCount', { v: t.materialCount })}} |
            {{$t('taskList.recordCount', { v: t.recordCount })}}</div>
          <div class="text-subtle">{{formatDate(t.startTime)}} ~ {{formatDate(t.endTime)}}</div>
      </div>
    </div>
  </div>

  <el-pagination v-if="tasks" class="card" layout="prev, pager, next, total" :total="tasks.total"
    :current-page="Number(params.page || 1)" :page-size="Number(params.size || 9)"
    @update:current-page="v => params.page = String(v)" @update:page-size="v => params.size = String(v)"
  ></el-pagination>

  <el-dialog v-model="isDialogOpen" :title="$t('project.task.dialog.title')">
    <task-editor v-model="form">
    </task-editor>
    <template #footer>
      <el-button
        type="primary"
        @click="request('POST', `/projects/${id}/tasks`, form).then(id => id && $router.push(`/tasks/${id}`))"
      >
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
