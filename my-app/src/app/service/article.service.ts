import { Injectable } from '@angular/core';
import { HttpService } from './http.service';

@Injectable({
  providedIn: 'root'
})
export class ArticleService {

  constructor(private http: HttpService) {
  }

  // 文章列表
  articleList = (data?: any) => {
    return this.http.request({
      method: 'post',
      url: '/article/articleList',
      data
    });
  }
  // 文章详情
  getDetail = (params: any) => {
    return this.http.request({
      method: 'get',
      url: '/article/articleDetail',
      params
    });
  }
  // 验证加密文章密码
  surePassword = (params: any) => {
    return this.http.request({
      method: 'get',
      url: '/article/articleDetail',
      params
    });
  }
  // 获取文章分类
  getArticleClassification = (params?: any) => {
    return this.http.request({
      method: 'get',
      url: '/article/articleClassify',
      params
    });
  }
  // 保存文章
  addArticle = (data: object) => {
    return this.http.request({
      method: 'post',
      url: '/adminArticle/add',
      data
    });
  }
  //  热门文章
  hotArticle = () => {
    return this.http.request({
      method: 'get',
      url: '/article/hotArticle',
    });
  }
  //  获取所有tag
  getTag = () => {
    return this.http.request({
      method: 'get',
      url: '/article/tag',
    });
  }
  // 获取配置
  getConfig = () => {
    return this.http.request({
      method: 'get',
      url: '/article/config',
    });
  }

}
