<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Application, ApplicationStatus, ApplicationType } from '@/utils/tables';
import { formatDate, shortcuts } from '@/utils/utils';
import { Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import { ref, watchEffect } from 'vue';

const applications = ref<{
  data: Application[],
  total: number
} | null>(null)

const params = useUrlSearchParams<Record<string, string>>('history')
watchEffect(async () => applications.value = await request('GET', '/applications', params))

const applicationStatuses = ref<ApplicationStatus[]>([])
request<ApplicationStatus[]>('GET', '/application-statuses').then(res => applicationStatuses.value = res)

const applicationTypes = ref<ApplicationType[]>([])
request<ApplicationType[]>('GET', '/application-types').then(res => applicationTypes.value = res)

</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div class="text-xl">
      {{$t('application')}}
    </div>
    <div class="flex gap-4 flex-wrap">
      <el-input v-model="params.title" class="!w-64" :prefix-icon="Search" :placeholder="$t('search', { v: $t('title') })">
      </el-input>
      <el-date-picker type="datetimerange"
        :model-value="params.time"
        @update:model-value="v => params.time = v?.map((e: Date) => e.toISOString())"
        clearable :shortcuts="shortcuts"
        :start-placeholder="$t('select', { v: $t('startTime') })"
        :end-placeholder="$t('select', { v: $t('endTime') })"
      ></el-date-picker>
    </div>
    <div class="flex gap-4 flex-wrap">
      <el-button @click="delete params.statusId">
        {{$t('applicationList.typeAll')}}
      </el-button>
      <el-segmented v-model="params.statusId" :options="applicationStatuses" :props="{ value: 'id' }">
        <template #default="{item}">
          {{item.name}} {{item.applicationCount}}
        </template>
      </el-segmented>
      <el-button @click="delete params.typeId">
        {{$t('applicationList.statusAll')}}
      </el-button>
      <el-segmented v-model="params.typeId" :options="applicationTypes" :props="{ value: 'id' }">
        <template #default="{item}">
          {{item.name}} {{item.applicationCount}}
        </template>
      </el-segmented>
    </div>
    <div v-if="applications" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="a in applications.data" :key="a.id" @click="$router.push(`/applications/${a.id}`)" class="card cursor-pointer flex flex-col gap-2">
        <div class="space-x-2 line-clamp-1">
          <tag :color="a.status.color">
            {{a.status.name}}
          </tag>
          <tag :color="a.type.color">
            {{a.type.name}}
          </tag>
          <span class="text-lg">{{a.title}}</span>
        </div>
        <div class="text-subtle line-clamp-2">{{a.content}}</div>
        <div class="text-subtle space-x-2">
          <span>{{$t('applicant')}}: {{a.applicant.name}}</span>
          <span v-if="a.approver">{{$t('approver')}}: {{a.approver.name}}</span>
        </div>
        <div class="mt-auto text-subtle">
          {{formatDate(a.createdAt)}}
        </div>
      </div>
    </div>
    <el-pagination v-if="applications" class="card" layout="prev, pager, next, total" :total="applications.total"
      :current-page="Number(params.page || 1)" :page-size="Number(params.size || 9)"
      @update:current-page="v => params.page = String(v)" @update:page-size="v => params.size = String(v)"
    ></el-pagination>
  </div>
</template>
