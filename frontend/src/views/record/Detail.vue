<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Record} from '@/utils/tables';
import {formatDate} from '@/utils/utils';
import {Back} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useRoute} from 'vue-router';

const route = useRoute()

const record = ref<Record | null>(null)
request<Record>('GET', `/records/${route.params.id}`).then(res => record.value = res)
</script>

<template>
  <div v-if="record" class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div class="card flex items-center gap-2 flex-wrap">
      <el-button @click="$router.back" :icon="Back">
        {{$t('back')}}
      </el-button>
      <div class="text-lg">{{$t('record')}} {{record.title}}</div>
      <div class="text-sm-subtle link" @click="$router.push(`/users/${record.user.id}`)">{{record.user.name}}</div>
    </div>
    <div class="flex flex-wrap gap-2">
      <div v-for="v, k in {
        createdAt: formatDate(record.createdAt),
        task: record.task.name,
        project: record.task.project.name,
        client: record.task.project.client.name,
        workHours: record.minutes,
        }" :key="k" class="space-x-2">
        <span class="text-sm-subtle">{{$t(k)}}</span>
        <span>{{v}}</span>
      </div>
    </div>
    <div class="bg-subtle p-4 rounded-2xl">
      {{record.content}}
    </div>
    <div class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <el-image v-for="i in ['a1.png', 'a2.png', 'a3.png']" :key="i" :src="`/${i}`" class="rounded-2xl">
      </el-image>
    </div>
  </div>
</template>
