DROP PROCEDURE IF EXISTS `proc_search_slug`;

DELIMITER $$

CREATE PROCEDURE `proc_search_slug`(
  IN path VARCHAR(255),
  IN last_path VARCHAR(30),
  OUT result VARCHAR(255)
)
BEGIN
  DECLARE cnt int;
  DECLARE idx int DEFAULT 1;
  DECLARE search VARCHAR(255);
  DECLARE saved VARCHAR(255);

  SET saved = path;
  SELECT CONCAT(path, '/') INTO search;

  SELECT CHAR_LENGTH(last_path) INTO @len;

  label1: LOOP
    IF idx > @len THEN
      LEAVE label1;
    END IF;

    SELECT CONCAT(search, SUBSTRING(last_path FROM idx FOR 1)) INTO search;
    SELECT CONCAT('%', search, '%') INTO @tmp;

    # Query
    SELECT count(gc.group_count) into cnt FROM
      (SELECT count(slug) AS group_count  FROM page_slug WHERE slug LIKE @tmp GROUP BY slug) AS gc;

    IF cnt > 0 THEN
      SET idx = idx + 1;
      SET saved = search;
      ITERATE label1;
    END IF;

    LEAVE label1;

  END LOOP label1;

  SET result = saved;

END $$

DELIMITER ;

# CALL proc_search_slug("/se/soedermanland", "stockholm", @result);
# SELECT @result;