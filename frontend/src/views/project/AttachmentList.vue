<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Attachment} from '@/utils/tables';
import {fileIconMap, type List} from '@/utils/utils';
import {Delete, Download, Upload} from '@element-plus/icons-vue';
import {useRouteParams} from '@vueuse/router';
import {ref, watchEffect} from 'vue';

const attachments = ref<List<Attachment> | null>(null)

const id = useRouteParams('id')

watchEffect(async () => attachments.value = await request('GET', `/projects/${id.value}/attachments`))

</script>

<template>

  <div class="text-xl flex justify-between">
    <div>{{$t('attachment')}}</div>
    <el-button :icon="Upload">{{$t('upload', { name: $t('attachment') })}}</el-button>
  </div>

  <div v-if="attachments" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="a in attachments.data" :key="a.id" class="card cursor-pointer space-y-2">
      <div class="flex items-center">
        <div class="flex gap-2">
          <Icon v-if="fileIconMap[a.filename.split('.')[1]]" :icon="fileIconMap[a.filename.split('.')[1]].icon" :class="[fileIconMap[a.filename.split('.')[1]].color, 'text-xl shrink-0']" />
          <span class="truncate">{{a.name}}</span>
        </div>
        <el-button :icon="Download" class="ms-auto" circle></el-button>
        <el-button :icon="Delete" circle></el-button>
      </div>
      <div class="text-subtle">{{a.filename}}</div>
    </div>
  </div>

</template>
