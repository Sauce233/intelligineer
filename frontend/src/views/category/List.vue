<script setup lang="ts">
import { request } from '@/utils/axios';
import type { Category } from '@/utils/tables';
import { Plus, Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import { reactive, ref, watchEffect } from 'vue';
import type {List} from '@/utils/utils';
import CategoryEditor, {type Form} from '@/components/editor/CategoryEditor.vue';
import {useI18n} from 'vue-i18n';

const categories = ref<List<Category> | null>(null)

const params = useUrlSearchParams<Record<string, string>>('history')

watchEffect(async () => categories.value = await request('GET', '/categories', params))

const isDialogOpen = ref(false)

const form = reactive<Form>({
  name: '', profile: '',
})

const { t } = useI18n({

})
</script>

<template>
  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div class="flex justify-between">
      <div class="text-xl">{{$t('category')}}</div>
      <el-button :icon="Plus" @click="isDialogOpen = true">
        {{$t('openDialog')}}
      </el-button>
    </div>

    <div class="flex gap-4">
      <el-input
        v-model="params.name"
        class="!w-64"
        :prefix-icon="Search"
        :placeholder="$t('search', { v: $t('category') })"
      />
    </div>

    <div v-if="categories" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div
        v-for="t in categories.data"
        :key="t.id"
        @click="$router.push(`/categories/${t.id}`)"
        class="card cursor-pointer flex flex-col gap-2"
      >
        <div class="text-lg font-bold">
          {{ t.name }}
        </div>

        <div class="text-gray-600 line-clamp-3">
          {{ t.profile }}
        </div>

        <div class="mt-auto flex justify-between text-subtle">
          <div>{{t.materials.map(m => m.name).join('，')}} 等等 {{$t('count', { k: $t('material'), v: t.materialCount })}}</div>
        </div>
      </div>
    </div>

    <el-pagination v-if="categories"
      class="card"
      layout="prev, pager, next, total"
      :total="categories.total"
      :current-page="Number(params.page || 1)"
      :page-size="Number(params.size || 9)"
      @update:current-page="v => params.page = String(v)"
      @update:page-size="v => params.size = String(v)"
    ></el-pagination>
  </div>

  <el-dialog v-model="isDialogOpen" :title="t('dialogTitle')">
    <category-editor v-model="form" />
    <template #footer>
      <el-button>
        $t('confirm')
      </el-button>
    </template>
  </el-dialog>

</template>
