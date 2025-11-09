<script setup lang="ts">
import {request} from '@/utils/axios';
import type {TaskMaterial} from '@/utils/tables';
import type {List} from '@/utils/utils';
import {Delete, Plus} from '@element-plus/icons-vue';
import {useUrlSearchParams} from '@vueuse/core';
import {useRouteParams} from '@vueuse/router';
import {ref, watchEffect} from 'vue';

const id = useRouteParams('id')
const params = useUrlSearchParams<any>('history')
const taskMaterials = ref<List<TaskMaterial> | null>(null)
watchEffect(async () => taskMaterials.value = await request('GET', `/tasks/${id.value}/materials`, params))
</script>

<template>

  <div class="text-xl flex justify-between">
    <div>{{$t('taskDetail.materialConsumption')}}</div>
    <el-button :icon="Plus">{{$t('add', { name: $t('material') })}}</el-button>
  </div>
  <div v-if="taskMaterials" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="m in taskMaterials.data" :key="m.material.id" class="card space-y-2">
      <div class="flex items-center">
        <div>{{m.material.name}}</div>
        <el-button :icon="Delete" circle class="!ms-auto"></el-button>
      </div>
      <div class="text-subtle">消耗: {{m.quantity}} {{m.material.unit}}</div>
      <div class="text-subtle line-clamp-2">备注: {{m.profile}}</div>
    </div>
  </div>
</template>
