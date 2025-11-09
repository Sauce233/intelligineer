<script setup lang="ts">
import {usePersistedStore} from '@/stores/persisted';
import {ArrowLeft} from '@element-plus/icons-vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n({ messages: {
  zh: {
    colorMode: '颜色模式',
    auto: '自动',
    dark: '暗色',
    light: '亮色',
    serverAddr: '服务器地址',
  },
} })

const languages = [
  { id: 'en', label: 'English' },
  { id: 'zh-CN', label: '简体中文' },
  { id: 'zh-TW', label: '繁體中文' },
  { id: 'ja', label: '日本語' },
  { id: 'ko', label: '한국어' },
  { id: 'fr', label: 'Français' },
  { id: 'de', label: 'Deutsch' },
  { id: 'es', label: 'Español' },
  { id: 'it', label: 'Italiano' },
  { id: 'pt', label: 'Português' },
  { id: 'ru', label: 'Русский' },
  { id: 'ar', label: 'العربية' },
]

const persisted = usePersistedStore()

</script>

<template>
  <div class="max-w-5xl mx-auto p-4 space-y-4">

    <div class="card flex items-center gap-4">
      <el-button @click="$router.back" :icon="ArrowLeft">
        {{$t('back')}}
      </el-button>
      <div class="text-lg">
        {{$t('settings.title')}}
      </div>
    </div>

    <div class="card space-y-4">
      <page-header :content="t('setting')" @back="$router.back">
      </page-header>
      <el-form label-width="auto" label-position="top">
        <el-form-item :label="t('colorMode')">
          <el-radio-group v-model="persisted.colorMode">
            <el-radio-button
              v-for="item in ['auto', 'dark', 'light']"
              :key="item"
              :label="t(item)"
              :value="item"
            ></el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="Language / 语言 / Langue / Idioma">
          <el-select v-model="$i18n.locale" class="w-48">
            <el-option
              v-for="language in languages"
              :key="language.id"
              :label="language.label"
              :value="language.id"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="t('serverAddr')">
          <el-input v-model="persisted.serverAddr" class="w-48">
          </el-input>
        </el-form-item>
      </el-form>
    </div>

  </div>
</template>
