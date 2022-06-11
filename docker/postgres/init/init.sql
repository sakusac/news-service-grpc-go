DROP TABLE IF EXISTS news;
CREATE TABLE news
(
    id SERIAL NOT NULL ,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO news (
    title,
    content,
    created_at,
    updated_at
) VALUES
    (
        'ジダンの「頭突き像」、Ｗ杯開催のカタールで再設置へ',
        '［ドーハ　６日　ロイター］ - サッカーの２００６年ワールドカップ（Ｗ杯）決勝でジネディーヌ・ジダンが相手選手に頭突きした姿を再現した銅像が、カタールで再び設置されることになった。同国の博物館・美術館を管理、運営するカタール・ミュージアムズが６日発表した。',
        '2022-06-07 10:59:00.000000',
        '2022-06-08 13:00:00.000000'
    ),
    (
        '米議会でＵＦＯ巡り半世紀ぶり公聴会、国防総省「解明に尽力」',
        '［ワシントン　１７日　ロイター］ - 米下院情報特別委員会の小委員会は１７日、未確認飛行物体（ＵＦＯ）に関する公聴会を約５０年ぶりに開いた。出席した国防総省高官は、ＵＦＯ現象の解明に尽力しているが、その多くがなお説明できないと語った。',
        '2022-05-18 12:40:00.000000',
        '2022-05-30 11:35:00.000000'
    );
