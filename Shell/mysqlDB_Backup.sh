#!/bin/bash

cd /home/devDBbackup

backupDirectory(){
        backup="/home/devDBbackup/dbbackup"

        if [ ! -d $backup ];then
                mkdir $backup
        fi

        cd $backup
        vy=`date +%Y`
        vd=`date +%m-%d`

        if [ -d $vy ];then
                echo "当前年目录已存在"
                cd "$vy"
                if [ -e $vd ];then
                        echo "当前年的日期目录已存在"
                        cd "$vd"
                        dbBackup
                        backupRetentionTime
                else
                        mkdir "$vd"
                        cd "$vd"
                        dbBackup
                        backupRetentionTime
                fi
        else
                mkdir "$vy"
                cd "$vy"
                mkdir "$vd"
                cd "$vd"
                dbBackup
                backupRetentionTime
        fi
}

dbBackup(){

        mkdir temporary
        mkdir suc
        mkdir fail
        cd temporary

        alldb=`mysql -uroot -pdosion123456 -e "show databases;" | grep -Ev "Database|information_schema|mysql|performance_schema|sys"`
        for i in $alldb;do
                mysqldump --databases $i > "$i".sql
                if [ $? -eq 0 ]; then
                        mv "$i".sql ../suc/
                        echo $i "备份成功" >> /home/devDBbackup/successBackup.log
                else
                        mv "$i".sql ../fail/
                        echo $i "备份失败" >> /home/devDBbackup/errBackup.log
                fi
        done

        echo "--------------------备份时间:`date +%Y-%m-%d`--------------------" >> /home/devDBbackup/successBackup.log
        echo "--------------------备份时间:`date +%Y-%m-%d`--------------------" >> /home/devDBbackup/errBackup.log

}

backupRetentionTime(){
        vy=`date +%Y`
        today=`date +%d`
        last_first_day=`echo $today | xargs -I{} date -d '+1 month -{} day' +%d`
        if [ $today -ge $last_first_day ]; then
                echo "Today is the last day of the month!"
                find /home/devDBbackup/dbbackup/$vy/* -ctime +30 | xargs rm -rf
                find /home/devDBbackup/dbbackup/$vy/* -ctime +30 | xargs rm -rf
        else
                echo "Today is not the last day of the month!"
        fi
}

backupDirectory


——————————————————————————————————————————————————————————


参考代码：
1、在每月最后二天执行指定命令
#!/bin/bash
#取得当前日期的天数部分
today=`date +%d`
#取得本月最后一天的天数部分
last_first_day=`echo $today | xargs -I{} date -d '+1 month -{} day' +%d`
#取得本月倒数第二天的天数部分
last_second_day=`echo $today | xargs -I{} date -d '+1 month -{} day -1 day' +%d`
#通过比较当前日期和倒数第二天的天数部分还决定是否执行指定命令
if [ $today -ge $last_second_day ]; then
    echo "Today is last two days of this month!"
else
    echo "Today is not last two days of this month!"
fi


2、递归遍历目录下的所有文件并删除5分钟前的文件
#!/bin/bash

jd_tim() {
    filename=$1

    timestamp=$(date +%s)
    filetimestamp=$(stat -c %Y $filename)


    if [ $[$timestamp - $filetimestamp] -lt 300 ]; then
        echo "less than five min ========== " $1
    else
        echo timestamp $timestamp
        echo filetimes $filetimestamp
        echo during is $[$timestamp - $filetimestamp]
        echo "delete it >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>delete " $1
        rm $1 -f
    fi

}

read_dir(){

    for file in `ls -a $1`
    do
        if [ -d $1"/"$file ]; then
            #echo in "-d file " $file
            if [[ $file != '.' && $file != '..' ]]; then
                #echo in "-------------------------------------- file " $file
                read_dir $1"/"$file
            #else
            #    echo "else iiii file is " $file
            fi

        else
            #echo "else-->>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "  $1"/"$file
            jd_tim $1"/"$file
        fi
    done
}



while [ 1 ]
do
    echo `date` " : " "$1 is " $1
    read_dir $1
    sleep 10
done

调用：
/bin/bash /usr/bin/judge.sh /mnt/face/faceCapture/