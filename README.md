

- 临时写的奖项批量格式化工具，用来把各种赛事奖项用正式名称包装起来，并对人生成奖项汇总
- 做了一些简单的映射规则，这样某某某某某某比赛一等奖就只用在表格某格写一个 1 了
- 用于当你突然需要把很多人在同几项比赛中获得的不同奖项用正式名称迅速整理出来时
- 能用，但也仅仅是能用而已
- ~~不知道除了我还有谁会用到~~

#### To Use

> contest sheet: 填写赛事简称，名称，不同奖项简写数字映射规则
> 
> data sheet: 填写选手姓名，奖项简写。同一比赛不同奖项用中文逗号隔开，识别词组内的数字，文字单独处理
> 
> contest+ sheet: 填写选手姓名，不在 contest sheet 里的奖项

示例运行结果（data.xlsx → output.txt）：

```txt
participant1：宇宙首届好味道泡面大赏二等奖，2021多喝热水高校小组赛参赛

participant2：第32届CPU大师区域赛冥王站铜牌，最烂小组作业全国赛团队三等奖，2022年度找猫猫比赛个人二等奖，团队二等奖

participant3：最烂小组作业全国赛团队一等级，2022年度找猫猫比赛个人三等奖，宇宙首届好味道泡面大赏三等奖，2021多喝热水高校小组赛参赛```
