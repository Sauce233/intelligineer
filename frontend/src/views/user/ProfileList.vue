<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Profile} from '@/utils/tables';
import type {List} from '@/utils/utils';
import {Download, Plus} from '@element-plus/icons-vue';
import {useUrlSearchParams} from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import {ref, watchEffect} from 'vue';

const params = useUrlSearchParams<any>('history')
const id = useRouteParams('id')
const profiles = ref<List<Profile> | null>(null)
watchEffect(async () => profiles.value = await request('GET', '/profiles', { userId: id.value }))

</script>

<template>

  <div class="flex justify-between">
    <div class="text-2xl">{{$t('userDetail.profileTitle')}}</div>
    <el-button :icon="Plus">{{$t('userDetail.addProfile')}}</el-button>
  </div>

  <div v-if="profiles" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="p in profiles.data" :key="p.id" class="card flex flex-col gap-2">
      <div class="flex items-center gap-2">
        <tag :color="p.type.color" class="shrink-0">{{p.type.name}}</tag>
        <span class="truncate">{{p.name}}</span>
      </div>
      <div class="text-subtle">{{p.profile}}</div>
      <div class="mt-auto">
        <span class="text-subtle">{{p.filename}}</span>
        <el-button :icon="Download">{{$t('download')}}</el-button>
      </div>
    </div>
  </div>

  <el-pagination
    v-if="profiles"
    layout="prev, pager, next, total"
    :total="profiles.total"
    :current-page="Number(params.page) || 1"
    @update:current-page="v => params.page = v"
    :page-size="Number(params.size) || 12"
    @update:page-size="v => params.size = v"
  />

</template>
