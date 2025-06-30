#!/bin/bash

echo "🔧 بدء عملية التنظيف ..."

# تحديث النظام
sudo apt update && sudo apt upgrade -y

# حذف الحزم اليتيمة
echo "🧹 حدف الحزم اليتيمة واللي ما بقاتش ضرورية..."
sudo apt autoremove --purge -y
sudo apt clean
sudo apt autoclean

# تصحيح الحزم المكسورة
echo "🧪 تصحيح الحزم المكسورة إن وجدت..."
sudo dpkg --configure -a
sudo apt -f install -y

# حذف الصور المصغرة (thumbnails)
echo "🖼️ حذف الصور المصغرة..."
rm -rf ~/.cache/thumbnails/*

# تنظيف سجلات النظام (logs) لأكثر من 7 أيام
echo "🧾 تنظيف السجلات القديمة..."
sudo journalctl --vacuum-time=7d

# حذف الملفات المكررة (اختياري - تأكد قبل الموافقة)
echo "🌀 هل ترغب في حذف الملفات المكررة؟ (y/n)"
read confirm
if [[ "$confirm" == "y" || "$confirm" == "Y" ]]; then
    echo "🔍 تثبيت fdupes ثم البحث عن الملفات المكررة..."
    sudo apt install -y fdupes
    fdupes -rdN /home/$USER
fi

echo "✅ التنظيف تم بنجاح!"
