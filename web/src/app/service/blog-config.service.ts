import { Injectable } from '@angular/core';
import { ArticleService } from './article.service';
import { toTimeDH } from '../../util/util';
@Injectable({
  providedIn: 'root'
})
export class BlogConfigService {
  configObservable: any;

  constructor(private httpService: ArticleService) { }
  public BlogName = '';
  public AuthorName = '';
  public AuthorMotto = '';
  public BlogLogo = '';
  public AuthorAvatar = '';
  public ArticleCount = '';
  public BlogTime = '';
  public ActiveTime = '';
  public BlogViewCount = '';
  public FilingMsg = '';
  public BlogNotice = '';

  getConfig = async () => {
    const res = await this.httpService.getConfig();
    if (res.code !== 200) { return; }
    this.BlogName = res.data.blog_name;
    this.AuthorName = res.data.author_name;
    this.AuthorMotto = res.data.author_motto;
    this.BlogLogo = res.data.blog_logo;
    this.AuthorAvatar = res.data.author_avatar;
    this.ArticleCount = res.data.article_count;
    this.BlogViewCount = res.data.blog_view_count;
    this.FilingMsg = res.data.filing_msg;
    // 处理时间
    this.BlogTime = toTimeDH(res.data.blog_start_time, 'DD-HH');
    this.ActiveTime = toTimeDH(res.data.active_time, 'DD');
    this.BlogNotice = res.data.blog_notice;
    this.BlogLogo = res.data.blog_logo;
  }
}