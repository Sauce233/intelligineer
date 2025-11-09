<script setup lang="ts">
import {request} from '@/utils/axios';
import type {Category} from '@/utils/tables';
import {Back} from '@element-plus/icons-vue';
import {useRouteParams} from '@vueuse/router';
import {ref} from 'vue';

const id = useRouteParams('id')
const category = ref<Category | null>(null)

request<Category>('GET', `/categories/${id.value}`).then(res => category.value = res)
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div v-if="category" class="card space-x-2">
      <el-button @click="$router.back" :icon="Back">
        {{$t('back')}}
      </el-button>
      <span class="text-lg">{{category.name}}</span>
    </div>

    <div v-if="category" class="bg-subtle p-4">
      {{category.profile}}
    </div>

  </div>
</template>
