<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Application } from '@/utils/tables';
import { type List } from '@/utils/utils';
import {useUrlSearchParams} from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import { ref, watchEffect } from 'vue';

const id = useRouteParams('id')
const params = useUrlSearchParams('history')
const applications = ref<List<Application> | null>(null)

watchEffect(async () => applications.value = await request('GET', '/applications', { applicantId: id.value }))

</script>

<template>
  <div class="text-2xl flex justify-between">
    <div>{{$t('userDetail.applicationsTitle')}}</div>
  </div>
  <div v-if="applications" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="a in applications.data" :key="a.id" class="card space-y-2">
      <div class="space-x-2">
        <tag :color="a.status.color">{{a.status.name}}</tag>
        <tag :color="a.type.color">{{a.type.name}}</tag>
        <span>{{a.title}}</span>
      </div>
      <div class="text-subtle line-clamp-2">{{a.content}}</div>
      <div v-if="a.approver" class="text-subtle">{{$t('approver')}} {{a.approver.name}}</div>
    </div>
  </div>

  <el-pagination v-if="applications" class="card" layout="prev, pager, next, total" :total="applications.total"
    :current-page="Number(params.page ?? 1)" @update:current-page="v => params.page = String(v)"
    :page-size="Number(params.size ?? 9)" @update:page-size="v => params.size = String(v)"
  ></el-pagination>
</template>
