<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import {request} from '@/utils/axios';
import type {UserAttendance} from '@/utils/tables';
import {formatDate, type List} from '@/utils/utils';
import {useUrlSearchParams} from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import {ref, watchEffect} from 'vue';

const params = useUrlSearchParams<any>('history')
const id = useRouteParams('id')
const userAttendances = ref<List<UserAttendance> | null>(null)
watchEffect(async () => userAttendances.value = await request('GET', `/users/${id.value}/attendances`, params))

</script>

<template>

  <div class="text-2xl">{{$t('userDetail.attendancesTitle')}}</div>

  <div v-if="userAttendances" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="a in userAttendances.data" class="card flex flex-col gap-2">
      <div class="flex gap-2 items-center">
        <tag :color="a.status.color" class="shrink-0">{{a.status.name}}</tag>
        <tag :color="a.attendance.type.color" class="shrink-0">{{a.attendance.type.name}}</tag>
        <span class="truncate">{{a.attendance.name}}</span>
      </div>
      <div v-if="a.time" class="text-subtle">{{formatDate(a.time)}}</div>
    </div>
  </div>

  <el-pagination
    v-if="userAttendances"
    layout="prev, pager, next, total"
    :total="userAttendances.total"
    :current-page="Number(params.page) || 1"
    @update:current-page="v => params.page = v"
    :page-size="Number(params.size) || 12"
    @update:page-size="v => params.size = v"
  />

</template>
