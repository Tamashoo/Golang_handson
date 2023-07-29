-- SQLite
CREATE TABLE "mydata" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL,
    "mail" TEXT,
    "age" INTEGER
);

INSERT INTO mydata (id, name, mail, age)
    values(1, 'Taro', 'taro@yamada', 39),
    (2, 'Hanako', 'hanako@flower', 28),
    (3, 'Sachiko', 'sachiko@happy', 17),
    (4, 'Jiro', 'jiro@change', 6);