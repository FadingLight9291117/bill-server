DROP DATABASE IF EXISTS bill;
CREATE DATABASE IF NOT EXISTS bill;
USE bill;

create table Bills
(
    id      int auto_increment primary key,
    type    varchar(15)  not null,
    date    varchar(15)  not null,
    money   varchar(900) null,
    cls     varchar(900) null,
    label   varchar(900) null,
    options varchar(900) null
);

INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3341, 'INCOME', '2022-09-01', '8.5', '交通', '打的', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3342, 'INCOME', '2022-09-01', '2.0', '交通', '公交', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3343, 'INCOME', '2022-09-01', '43.0', '交通', '其他', '巴士 回房县');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3344, 'INCOME', '2022-09-01', '5.0', '餐饮', '饮料', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3345, 'INCOME', '2022-09-01', '4.0', '交通', '公交', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3346, 'INCOME', '2022-09-04', '168.0', '餐饮', '聚餐', '胖嫂烧烤');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3347, 'INCOME', '2022-09-03', '27.0', '餐饮', '零食', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3348, 'INCOME', '2022-09-06', '186.0', '交通', '高铁', '回武汉的高铁票');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3349, 'INCOME', '2022-09-06', '4.0', '餐饮', '饮料', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3350, 'INCOME', '2022-09-06', '5.0', '餐饮', '早餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3351, 'INCOME', '2022-07-04', '17.9', '餐饮', '晚餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3352, 'INCOME', '2022-07-04', '18.5', '餐饮', '午餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3353, 'INCOME', '2022-07-04', '9.0', '餐饮', '饮料', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3354, 'INCOME', '2022-07-04', '17.0', '餐饮', '早餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3355, 'INCOME', '2022-07-04', '10.3', '餐饮', '零食', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3356, 'INCOME', '2022-07-05', '9.5', '餐饮', '早餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3357, 'INCOME', '2022-07-05', '22.0', '餐饮', '午餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3358, 'INCOME', '2022-07-05', '16.0', '餐饮', '晚餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3359, 'INCOME', '2022-07-05', '1.6', '交通', '公交', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3360, 'INCOME', '2022-07-06', '7.0', '餐饮', '早餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3361, 'INCOME', '2022-07-06', '19.0', '餐饮', '午餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3362, 'INCOME', '2022-07-06', '6.0', '餐饮', '饮料', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3363, 'INCOME', '2022-07-06', '14.0', '餐饮', '晚餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3364, 'INCOME', '2022-07-06', '1.6', '交通', '公交', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3365, 'INCOME', '2022-07-06', '1.6', '交通', '公交', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3366, 'INCOME', '2022-07-06', '16.0', '餐饮', '奶茶', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3367, 'INCOME', '2022-07-07', '9.0', '餐饮', '早餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3368, 'INCOME', '2022-07-07', '14.4', '餐饮', '午餐', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3369, 'INCOME', '2022-07-07', '1.6', '交通', '公交', '');
INSERT INTO bill.Bills (id, type, date, money, cls, label, options)
VALUES (3370, 'INCOME', '2022-07-07', '9.0', '餐饮', '晚餐', '');