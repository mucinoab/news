CREATE TABLE NEWSLETTER (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    NAME TEXT NOT NULL,
    DESCRIPTION TEXT,
    SUBJECT TEXT NOT NULL,
    CONTENT_ATTACHMENT_PATH TEXT NOT NULL
);

CREATE TABLE RECEIPIENT (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    EMAIL TEXT NOT NULL,
    NEWSLETTER_ID INTEGER NOT NULL,

    FOREIGN KEY (NEWSLETTER_ID)
      REFERENCES NEWSLETTER(ID)
);

CREATE INDEX IDX_RECEIPIENT_NEWSLETTER_ID ON RECEIPIENT (NEWSLETTER_ID);
CREATE INDEX IDX_NEWSLETTER_ID_AND_RECEIPIENT_EMAIL ON RECEIPIENT (NEWSLETTER_ID, EMAIL);

INSERT INTO NEWSLETTER (NAME, DESCRIPTION, SUBJECT, CONTENT_ATTACHMENT_PATH)
VALUES ('NEWSLETTER Uno', 'DESCRIPTION Bla Bla Bla', 'A NICE SUBJECT', 'NOTHING_TO_SEE_HERE');

INSERT INTO RECEIPIENT (EMAIL, NEWSLETTER_ID)
VALUES ('MUCINOAB@GMAIL.COM', 1);
