WITH _overall_aggregation AS (
  SELECT
    "userName",
    COUNT(threatScore) AS threatScore_count
  from
    perf.IDM_SiteVisitArray
  WHERE
    (
      tenantid = '002fb12e-f185-496a-8b49-636b877a1358'
    )
    AND (
      eventTime BETWEEN '2024-02-01 23:19:00'
      AND '2024-02-21 23:19:00'
    )
  GROUP BY
    userName
),
userName_selector AS (
  SELECT
    userName AS userName_filtered,
    SUM(threatScore_count) as threatScore_count_userName,
    ROW_NUMBER() over (
      ORDER BY
        threatScore_count_userName DESC,
        userNam e ASC
    ) as userName_internal_rank
  FROM
    _overall_aggregation
  GROUP BY
    userName
)
SELECT
  "userName",
  threatScore_count,
  count(*) over() as fp_internal_count
from
  _overall_aggregation,
  userName_selector
WHERE
  (
    (
      userName = userName_selector.userName_filtered
      OR (
        userName IS NULL
        AND userName_selector.userName_filtered IS NULL
      )
    )
    AND userName_internal_rank <= 10
  )
LIMIT
  100