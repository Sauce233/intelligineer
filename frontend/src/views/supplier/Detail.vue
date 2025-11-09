<script setup lang="ts">
import type {Form} from '@/components/editor/SupplierEditor.vue';
import SupplierEditor from '@/components/editor/SupplierEditor.vue';
import {request} from '@/utils/axios';
import type {Supplier} from '@/utils/tables';
import {formatDate} from '@/utils/utils';
import {Back, Delete, Edit} from '@element-plus/icons-vue';
import {useRouteParams} from '@vueuse/router';
import {reactive, ref} from 'vue';
import {useI18n} from 'vue-i18n';

const id = useRouteParams('id')

const supplier = ref<Supplier | null>(null)

const load = () => request<Supplier>('GET', `/suppliers/${id.value}`).then(res => supplier.value = res)
load()

const isDialogOpen = ref(false)

const form = reactive<Form>({
  name: '', address: '', profile: '',
})

const { t } = useI18n({})

</script>

<template>
  <div v-if="supplier" class="max-w-5xl mx-auto flex flex-col gap-4 p-4">

    <div class="card flex gap-2 items-center flex-wrap">
      <el-button @click="$router.back" :icon="Back">
        {{$t('back')}}
      </el-button>
      <div>{{supplier.name}}</div>
      <el-button class="ms-auto" :icon="Edit" circle @click="isDialogOpen = true; Object.assign(form, supplier)" />
      <el-button :icon="Delete" circle @click="request('DELETE', `/suppliers/${id}`).then(ok => ok && $router.push(`/suppliers`))" />
    </div>

    <div class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div
        v-for="v in [
          ['name', supplier.name, ''],
          ['createdAt', formatDate(supplier.createdAt), ''],
          ['updatedAt', formatDate(supplier.updatedAt), ''],
          ['address', supplier.address, 'col-span-2'],
        ]"
        :key="v[0]"
        :class="['flex items-center gap-2', v[2]]"
      >
        <div class="text-subtle">{{$t(v[0])}}</div>
        <div>{{v[1]}}</div>
      </div>
    </div>

    <div class="text-xl">
      {{$t('profile')}}
    </div>

    <div class="bg-subtle p-4 rounded-2xl">
      {{supplier.profile}}
    </div>

    <div class="p-2 flex gap-1">
      <router-link
        v-for="v in [
          [`/suppliers/${id}`, $t('contact')],
          [`/suppliers/${id}/purchases`, $t('purchase')],
        ]"
        :key="v[0]"
        :to="v[0]"
        class="p-2 rounded-2xl"
        exact-active-class="font-bold bg-subtle"
      >
        {{v[1]}}
      </router-link>
    </div>

    <router-view />

  </div>

  <el-dialog v-model="isDialogOpen" :title="t('dialogTitle')">
    <supplier-editor v-model="form" />
    <template #footer>
      <el-button @click="request('PATCH', `/suppliers/${id}`, form).then(ok => ok && load())">
        {{t('confirm')}}
      </el-button>
    </template>
  </el-dialog>
</template>
