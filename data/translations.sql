# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 127.0.0.1 (MySQL 5.6.25)
# Database: device_management
# Generation Time: 2015-08-20 10:17:14 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table translations
# ------------------------------------------------------------

DROP TABLE IF EXISTS `translations`;

CREATE TABLE `translations` (
  `locale` varchar(12) DEFAULT NULL,
  `key` longtext,
  `value` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `translations` WRITE;
/*!40000 ALTER TABLE `translations` DISABLE KEYS */;

INSERT INTO `translations` (`locale`, `key`, `value`)
VALUES
	('en-US','qor_admin.设备管理','设备管理'),
	('en-US','qor_admin.Devices','设备维护'),
	('en-US','qor_admin.CustomerDeviceIncomings','客户设备收入'),
	('en-US','qor_admin.CustomerDeviceOutcomings','客户设备还回'),
	('en-US','qor_admin.Consumables','消耗品'),
	('en-US','qor_admin.查询','查询'),
	('en-US','qor_admin.ReportItems','设备当前状态'),
	('en-US','qor_admin.人事管理','人事管理'),
	('en-US','qor_admin.Clients','客户维护'),
	('en-US','qor_admin.Employees','员工维护'),
	('en-US','qor_admin.消耗品管理','消耗品管理'),
	('en-US','qor_admin.系统设置','系统设置'),
	('en-US','qor_admin.Translations','翻译'),
	('en-US','qor_admin.consumables.attributes.Code','编码'),
	('en-US','qor_admin.Actions','操作'),
	('en-US','qor_admin.Source','数据源'),
	('en-US','qor_admin.Target','目标'),
	('en-US','qor_admin.Bulk Edit','批量修改'),
	('en-US','qor_admin.Exit Bulk Edit','退出批量修改'),
	('en-US','qor_admin.Copy All','复制全部'),
	('en-US','qor_admin.Copy','复制'),
	('en-US','qor_admin.Translation','翻译'),
	('en-US','qor_admin.Saved','已经保存'),
	('en-US','qor_admin.Cancel Edit','取消修改'),
	('en-US','qor_admin.Save','保存'),
	('en-US','qor_admin.Dashboard','主页'),
	('en-US','qor_admin.Are you sure?','你确定要执行吗？'),
	('en-US','qor_admin.Show All','显示全部'),
	('en-US','qor_admin.qor_admin.report_items.ReportItems','设备状态查询'),
	('en-US','qor_admin.report_items.attributes.Who Has Them','当前位于'),
	('en-US','qor_admin.report_items.attributes.Device Name','设备名'),
	('en-US','qor_admin.report_items.attributes.Device Code','设备编码'),
	('en-US','qor_admin.report_items.attributes.Count','设备数量'),
	('en-US','qor_admin.customer_device_incomings.attributes.客户名','客户名'),
	('en-US','qor_admin.customer_device_incomings.attributes.设备名','设备名'),
	('en-US','qor_admin.customer_device_outcomings.attributes.客户名','客户名'),
	('en-US','qor_admin.customer_device_outcomings.attributes.设备名','设备名'),
	('en-US','qor_admin.日常操作','日常操作'),
	('en-US','qor_admin.数据维护','数据维护'),
	('en-US','qor_admin.Cancel','取消'),
	('en-US','qor_admin.Save Changes','保存修改'),
	('en-US','qor_admin.report_items.attributes.Client Name','客户'),
	('en-US','qor_admin.customer_device_incomings.attributes.仓库','仓库'),
	('en-US','qor_admin.customer_device_outcomings.attributes.仓库','仓库'),
	('en-US','qor_admin.DeviceIns','自有设备还回'),
	('en-US','qor_admin.DeviceOuts','自有设备带出'),
	('en-US','qor_admin.consumables.attributes.设备名','设备名'),
	('en-US','qor_admin.consumables.attributes.设备代码','设备代码'),
	('en-US','qor_admin.consumables.attributes.设备数量','设备数量'),
	('en-US','qor_admin.qor_admin.customer_device_incomings.客户名','客户名'),
	('en-US','qor_admin.qor_admin.customer_device_incomings.设备名','设备名'),
	('en-US','qor_admin.qor_admin.customer_device_incomings.仓库','仓库'),
	('en-US','qor_admin.ClientDeviceIns','客户设备收入'),
	('en-US','qor_admin.ClientDeviceOuts','客户设备还回'),
	('en-US','qor_admin.ConsumableIns','消耗品购买'),
	('en-US','qor_admin.client_device_ins.attributes.客户名','客户名'),
	('en-US','qor_admin.client_device_ins.attributes.设备名','设备名'),
	('en-US','qor_admin.client_device_ins.attributes.仓库','仓库'),
	('en-US','qor_admin.client_device_ins.attributes.Quantity','设备数量'),
	('en-US','qor_admin.client_device_ins.attributes.Date','收入日期'),
	('en-US','qor_admin.client_device_outs.attributes.客户名','客户名'),
	('en-US','qor_admin.client_device_outs.attributes.设备名','设备名'),
	('en-US','qor_admin.client_device_outs.attributes.仓库','仓库'),
	('en-US','qor_admin.client_device_outs.attributes.Date','日期'),
	('en-US','qor_admin.ConsumableOuts','消耗品使用'),
	('en-US','qor_admin.client_device_outs.attributes.取出仓库','取出仓库'),
	('en-US','qor_admin.client_device_ins.attributes.存入仓库','存入仓库'),
	('en-US','qor_admin.qor_admin.client_device_ins.Add ClientDeviceIn','收入客户设备'),
	('en-US','qor_admin.qor_admin.client_device_ins.Edit ClientDeviceIn','修改收入客户设备记录'),
	('en-US','qor_admin.qor_admin.client_device_ins.客户名','客户名'),
	('en-US','qor_admin.qor_admin.client_device_ins.设备名','设备名'),
	('en-US','qor_admin.qor_admin.client_device_ins.存入仓库','存入仓库'),
	('en-US','qor_admin.consumable_ins.attributes.设备名','设备名'),
	('en-US','qor_admin.consumable_ins.attributes.设备代码','设备代码'),
	('en-US','qor_admin.consumable_ins.attributes.设备数量','设备数量'),
	('en-US','qor_admin.consumable_outs.attributes.设备名','设备名'),
	('en-US','qor_admin.consumable_outs.attributes.设备代码','设备代码'),
	('en-US','qor_admin.consumable_outs.attributes.设备数量','设备数量'),
	('en-US','qor_admin.Warehouses','仓库维护'),
	('en-US','qor_admin.client_device_ins.attributes.Warehouse','收入到仓库'),
	('en-US','qor_admin.client_device_ins.attributes.Client Name','收入客户名'),
	('en-US','qor_admin.client_device_ins.attributes.Device Name','收入设备名'),
	('en-US','qor_admin.client_device_ins.attributes.By Whom','操作员'),
	('en-US','qor_admin.不能为空','不能为空'),
	('en-US','qor_admin.收入设备名不能为空','收入设备名不能为空'),
	('en-US','qor_admin.收入客户名不能为空','收入客户名不能为空'),
	('en-US','qor_admin.收入到的仓库不能为空','收入到的仓库不能为空'),
	('en-US','qor_admin.收入设备的数量要大于0','收入设备的数量要大于0'),
	('en-US','qor_admin.请选择操作员','请选择操作员'),
	('en-US','qor_admin.report_items.attributes.Operated By Whom','操作员'),
	('en-US','qor_admin.qor_admin.report_items.Who Has Them','当前位于'),
	('en-US','qor_admin.qor_admin.report_items.Client Name','客户'),
	('en-US','qor_admin.qor_admin.report_items.Device Name','设备名'),
	('en-US','qor_admin.qor_admin.report_items.Device Code','设备编码'),
	('en-US','qor_admin.qor_admin.report_items.Operated By Whom','操作员'),
	('en-US','qor_admin.qor_admin.report_items.Count','设备数量'),
	('en-US','qor_admin.client_device_outs.attributes.By Whom','操作员'),
	('en-US','qor_admin.client_device_outs.attributes.Client Device In ID','客户设备'),
	('en-US','qor_admin.选择收入过的客户设备','选择收入过的客户设备'),
	('en-US','qor_admin.report_items.attributes.Who Has Them Name','设备位于'),
	('en-US','qor_admin.带出设备不能为空','带出设备不能为空'),
	('en-US','qor_admin.设备带出人不能为空','设备带出人不能为空'),
	('en-US','qor_admin.带出设备的数量要大于0','带出设备的数量要大于0'),
	('en-US','qor_admin.Logged in as {{$1}}','{{$1}}'),
	('en-US','qor_admin.logo','<a href=\"{{.Prefix}}\"><span class=\"visuallyhidden\">QOR</span></a><a href=\"/\" target=\"_blank\">View Site <i class=\"material-icons md-14\" aria-hidden=\"true\">open_in_new</i></a>'),
	('en-US','qor_admin.powered_by','Powered by <a href=\"http://getqor.com\" target=\"_blank\">QOR</a>'),
	('en-US','qor_admin.Search','Search'),
	('en-US','qor_admin.qor_admin.client_device_ins.ClientDeviceIns','客户设备收入'),
	('en-US','qor_admin.qor_admin.client_device_outs.ClientDeviceOuts','客户设备还回'),
	('en-US','qor_admin.client_device_outs.attributes.Client Name','客户名'),
	('en-US','qor_admin.client_device_outs.attributes.Device Name','设备'),
	('en-US','qor_admin.client_device_outs.attributes.Quantity','设备数量'),
	('en-US','qor_admin.client_device_outs.attributes.Warehouse Name','仓库'),
	('en-US','qor_admin.qor_admin.devices.Devices','设备列表'),
	('en-US','qor_admin.devices.attributes.ID','ID'),
	('en-US','qor_admin.devices.attributes.Name','名称'),
	('en-US','qor_admin.devices.attributes.Code','编码'),
	('en-US','qor_admin.devices.attributes.Total Quantity','总数量'),
	('en-US','qor_admin.devices.attributes.Warehouse ID','仓库'),
	('en-US','qor_admin.devices.attributes.Category ID','分类'),
	('en-US','qor_admin.qor_admin.devices.ID','ID'),
	('en-US','qor_admin.qor_admin.devices.Name','名称'),
	('en-US','qor_admin.qor_admin.devices.Code','编码'),
	('en-US','qor_admin.qor_admin.devices.Total Quantity','总数'),
	('en-US','qor_admin.qor_admin.devices.Warehouse ID','仓库'),
	('en-US','qor_admin.qor_admin.devices.Category ID','分类'),
	('en-US','qor_admin.qor_admin.employees.Employees','员工列表'),
	('en-US','qor_admin.employees.attributes.ID','ID'),
	('en-US','qor_admin.employees.attributes.Name','姓名'),
	('en-US','qor_admin.employees.attributes.Mobile','手机号码'),
	('en-US','qor_admin.qor_admin.employees.ID','ID'),
	('en-US','qor_admin.qor_admin.employees.Name','姓名'),
	('en-US','qor_admin.qor_admin.employees.Mobile','手机号码'),
	('en-US','qor_admin.resource_successfully_created','{{.Name}} was successfully created'),
	('en-US','qor_admin.qor_admin.client_device_ins.Client Name','Client Name'),
	('en-US','qor_admin.qor_admin.client_device_ins.Device Name','Device Name'),
	('en-US','qor_admin.qor_admin.client_device_ins.Warehouse','Warehouse'),
	('en-US','qor_admin.qor_admin.client_device_ins.Quantity','Quantity'),
	('en-US','qor_admin.qor_admin.client_device_ins.By Whom','By Whom'),
	('en-US','qor_admin.qor_admin.client_device_ins.Date','Date'),
	('en-US','qor_admin.qor_admin.report_items.Who Has Them Name','Who Has Them Name'),
	('en-US','qor_admin.qor_admin.device_outs.DeviceOuts','自有设备带出'),
	('en-US','qor_admin.device_outs.attributes.To Whom Name','带出人'),
	('en-US','qor_admin.device_outs.attributes.From Warehouse Name','提取仓库'),
	('en-US','qor_admin.device_outs.attributes.Device Name','设备'),
	('en-US','qor_admin.device_outs.attributes.Quantity','数量'),
	('en-US','qor_admin.device_outs.attributes.By Whom Name','操作员'),
	('en-US','qor_admin.device_outs.attributes.Date','日期'),
	('en-US','qor_admin.qor_admin.device_outs.Add DeviceOut','设备带出'),
	('en-US','qor_admin.device_outs.attributes.From Report Item ID','当前设备'),
	('en-US','qor_admin.device_outs.attributes.To Whom ID','带出人'),
	('en-US','qor_admin.device_outs.attributes.By Whom ID','操作员'),
	('en-US','qor_admin.qor_admin.devices.Edit Device','编辑设备'),
	('en-US','qor_admin.qor_admin.devices.Add Device','添加设备'),
	('en-US','qor_admin.数量输入有误，不能大于10','数量输入有误，不能大于10'),
	('en-US','qor_admin.qor_admin.device_outs.Edit DeviceOut','Edit DeviceOut'),
	('en-US','qor_admin.qor_admin.device_outs.To Whom Name','带出人'),
	('en-US','qor_admin.qor_admin.device_outs.From Warehouse Name','提取仓库'),
	('en-US','qor_admin.qor_admin.device_outs.Device Name','设备'),
	('en-US','qor_admin.qor_admin.device_outs.Quantity','数量'),
	('en-US','qor_admin.qor_admin.device_outs.By Whom Name','操作员'),
	('en-US','qor_admin.qor_admin.device_outs.Date','日期'),
	('en-US','qor_admin.qor_admin.client_device_outs.Add ClientDeviceOut','还回客户设备'),
	('en-US','qor_admin.qor_admin.device_ins.DeviceIns','自有设备还回'),
	('en-US','qor_admin.device_ins.attributes.ID','ID'),
	('en-US','qor_admin.device_ins.attributes.From Report Item ID','带出去的设备'),
	('en-US','qor_admin.device_ins.attributes.From Whom Name','带出人'),
	('en-US','qor_admin.device_ins.attributes.Device Name','设备'),
	('en-US','qor_admin.device_ins.attributes.Quantity','数量'),
	('en-US','qor_admin.device_ins.attributes.To Warehouse ID','还入仓库'),
	('en-US','qor_admin.device_ins.attributes.To Warehouse Name','还入仓库'),
	('en-US','qor_admin.device_ins.attributes.By Whom ID','操作员'),
	('en-US','qor_admin.device_ins.attributes.By Whom Name','操作员'),
	('en-US','qor_admin.device_ins.attributes.Date','日期'),
	('en-US','qor_admin.qor_admin.device_ins.Add DeviceIn','还入自有设备'),
	('en-US','qor_admin.qor_admin.device_ins.Edit DeviceIn','Edit DeviceIn'),
	('en-US','qor_admin.qor_admin.device_ins.From Whom Name','带出人'),
	('en-US','qor_admin.qor_admin.device_ins.To Warehouse Name','还入仓库'),
	('en-US','qor_admin.qor_admin.device_ins.Device Name','设备'),
	('en-US','qor_admin.qor_admin.device_ins.Quantity','数量'),
	('en-US','qor_admin.qor_admin.device_ins.By Whom Name','操作员'),
	('en-US','qor_admin.qor_admin.device_ins.Date','日期'),
	('en-US','qor_admin.qor_admin.consumable_outs.ConsumableOuts','消耗品使用'),
	('en-US','qor_admin.consumable_outs.attributes.ID','ID'),
	('en-US','qor_admin.consumable_outs.attributes.Name','Name'),
	('en-US','qor_admin.consumable_outs.attributes.Code','编码'),
	('en-US','qor_admin.consumable_outs.attributes.Count','Count'),
	('en-US','qor_admin.consumable_outs.attributes.Device Name','设备'),
	('en-US','qor_admin.consumable_outs.attributes.Quantity','数量'),
	('en-US','qor_admin.consumable_outs.attributes.To Whom Name','使用人'),
	('en-US','qor_admin.consumable_outs.attributes.Warehouse Name','仓库'),
	('en-US','qor_admin.consumable_outs.attributes.By Whom Name','操作员'),
	('en-US','qor_admin.consumable_outs.attributes.Date','日期'),
	('en-US','qor_admin.qor_admin.consumable_outs.Add ConsumableOut','使用消耗品'),
	('en-US','qor_admin.consumable_outs.attributes.Report Item ID','当前消耗品'),
	('en-US','qor_admin.consumable_outs.attributes.To Whom ID','使用人'),
	('en-US','qor_admin.consumable_outs.attributes.By Whom ID','操作员'),
	('en-US','qor_admin.qor_admin.consumable_outs.Edit ConsumableOut','Edit ConsumableOut'),
	('en-US','qor_admin.qor_admin.consumable_outs.Device Name','消耗品名'),
	('en-US','qor_admin.qor_admin.consumable_outs.Quantity','数量'),
	('en-US','qor_admin.qor_admin.consumable_outs.To Whom Name','使用人'),
	('en-US','qor_admin.qor_admin.consumable_outs.Warehouse Name','仓库'),
	('en-US','qor_admin.qor_admin.consumable_outs.By Whom Name','操作员'),
	('en-US','qor_admin.qor_admin.consumable_outs.Date','日期'),
	('en-US','qor_admin.qor_admin.consumable_ins.ConsumableIns','消耗品购买'),
	('en-US','qor_admin.consumable_ins.attributes.Device Name','设备'),
	('en-US','qor_admin.consumable_ins.attributes.Quantity','数量'),
	('en-US','qor_admin.consumable_ins.attributes.Warehouse Name','仓库'),
	('en-US','qor_admin.consumable_ins.attributes.By Whom Name','操作员'),
	('en-US','qor_admin.consumable_ins.attributes.Date','日期'),
	('en-US','qor_admin.qor_admin.consumable_ins.Add ConsumableIn','购买消耗品'),
	('en-US','qor_admin.consumable_ins.attributes.Report Item ID','现有消耗品'),
	('en-US','qor_admin.consumable_ins.attributes.By Whom ID','操作员'),
	('en-US','qor_admin.qor_admin.consumable_ins.Edit ConsumableIn','Edit ConsumableIn'),
	('en-US','qor_admin.qor_admin.consumable_ins.Device Name','设备'),
	('en-US','qor_admin.qor_admin.consumable_ins.Quantity','数量'),
	('en-US','qor_admin.qor_admin.consumable_ins.Warehouse Name','仓库'),
	('en-US','qor_admin.qor_admin.consumable_ins.By Whom Name','操作员'),
	('en-US','qor_admin.qor_admin.consumable_ins.Date','日期'),
	('en-US','qor_admin.zh_CN','zh_CN'),
	('en-US','qor_admin.qor_admin.client_device_outs.Edit ClientDeviceOut','Edit ClientDeviceOut'),
	('en-US','qor_admin.qor_admin.client_device_outs.Client Name','Client Name'),
	('en-US','qor_admin.qor_admin.client_device_outs.Device Name','Device Name'),
	('en-US','qor_admin.qor_admin.client_device_outs.Quantity','Quantity'),
	('en-US','qor_admin.qor_admin.client_device_outs.Warehouse Name','Warehouse Name'),
	('en-US','qor_admin.qor_admin.client_device_outs.By Whom','By Whom'),
	('en-US','qor_admin.qor_admin.client_device_outs.Date','Date'),
	('en-US','qor_admin.qor_admin.warehouses.Warehouses','仓库列表'),
	('en-US','qor_admin.warehouses.attributes.ID','ID'),
	('en-US','qor_admin.warehouses.attributes.Name','名称'),
	('en-US','qor_admin.warehouses.attributes.Address','地址'),
	('en-US','qor_admin.qor_admin.warehouses.ID','ID'),
	('en-US','qor_admin.qor_admin.warehouses.Name','名称'),
	('en-US','qor_admin.qor_admin.warehouses.Address','地址'),
	('en-US','qor_admin.qor_admin.employees.Add Employee','添加员工'),
	('en-US','qor_admin.qor_admin.employees.Edit Employee','编辑员工'),
	('en-US','qor_admin.qor_admin.warehouses.Add Warehouse','添加仓库'),
	('en-US','qor_admin.Next Page','下一页'),
	('en-US','qor_admin.resource_successfully_updated','{{.Name}} was successfully updated'),
	('en-US','qor_admin.代码已经存在了，不能重复','代码已经存在了，不能重复'),
	('en-US','qor_admin.更新后的库存数量不能小于零','更新后的库存数量不能小于零'),
	('en-US','qor_admin.更新后的库存数量不能小于零，在库0','更新后的库存数量不能小于零，在库0'),
	('en-US','qor_admin.Select an Option','请选择一项'),
	('en-US','qor_admin.qor_admin.warehouses.Edit Warehouse','编辑仓库');

/*!40000 ALTER TABLE `translations` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;