# Spec: Seat Forecast

Last updated: 2026-11-13

## Problem and user

The user is a Rutgers student who needs a seat in a closed section during registration or add/drop. Today they cannot tell whether waiting is worth it. They either refresh for days and risk missing the course they need, or give up and take a backup too early.

Rutgers already has many course snipers: RUReg, RU Snipe, rutgers.quest, Schedru, SwiftRU and Discord bots. They all answer "tell me the moment it opens." None answers "how likely is it to open before add/drop ends, and when?" Seat Forecast answers that question, so a student can decide to wait or to take a backup.

## Success metrics

These are targets, not results. Each one is measured from data the system already stores.

| Metric | Target | How it is measured |
| --- | --- | --- |
| Sections polled per cycle | All sections in the current term's Schedule of Classes | Count of sections in each poll run, logged by the poller |
| Poller uptime | 99% over any 7-day window | Heartbeat row written each cycle; gaps counted from the table |
| Open/close events logged | Every observed transition, with no silent drops | Row count in the events table vs. diff of consecutive snapshots |
| p95 API latency | Under 300 ms for a section lookup | Request timing logged by the API, p95 computed over 7 days |
| Students using it in spring registration | To be set after launch traffic is known | Distinct lookups per day from API logs, no accounts needed |
| Forecast calibration | To be set after first full cycle | Predicted probability vs. observed open rate, bucketed |

## V1 features

V1 ships five features, all read-only for the student.

1. **Event logger:** The poller records every section open and close transition with a timestamp.
2. **Open probability:** A section lookup returns the chance it opens before add/drop ends.
3. **Section history:** Each section shows its past open/close events as a simple timeline.
4. **Backup hint:** The page shows the probability beside the section's status so a student can compare a wanted section with a backup.
5. **Public lookup page:** A plain web page takes a course or index number and shows the result, with no login.

## Out of scope

V1 cuts everything that is not a forecast.

- **Auto-registration and WebReg login:** Logging into WebReg on a student's behalf is against Rutgers policy. The product alerts and predicts only.
- **Push and SMS alerts:** Existing snipers already do this well. Competing on alert speed adds nothing.
- **Mobile app:** A web page covers the lookup. A native app costs weeks and adds no new decision support.
- **Other universities:** One school keeps the data model and the polling target simple.
- **Fancy ML:** A baseline from historical open rates is enough until a full cycle of data exists.

## Architecture and stack

The system is one Go service, one database and one static page, with no queue in v1.

| Component | Tech | Job |
| --- | --- | --- |
| Poller | Go | Reads public Schedule of Classes data on an interval and diffs section status |
| Event store | Postgres | Holds one row per open/close transition, plus poll heartbeats |
| Model | Go or small Python script | Computes open probability from historical rates by course, section and time in cycle |
| API | Go | Serves per-section probability and event history as JSON |
| Web page | Plain HTML and JS | Takes a section and shows probability and history |
| CI and deploy | Docker, GitHub Actions, small VM or free-tier host | Builds and tests on push; runs the poller and API unattended |

Data flows from the poller to Postgres as events, the model reads those events, and the API serves its output to the web page.

## Milestones

Two build windows ship the logger first and the forecast second, and a third step retrains on real data.

| Days | Work | Done when |
| --- | --- | --- |
| 2-3 | 1. Inspect the public Schedule of Classes data and define a section snapshot. 2. Write the Go poller and the diff that detects open/close changes. 3. Create the Postgres schema for events and heartbeats. 4. Dockerize and deploy to a VM or free-tier host. 5. Add GitHub Actions CI for build and tests. | It is logging open/close events to Postgres on a server, unattended. |
| 17-18 | 1. Export about 2 weeks of events and check data quality. 2. Build the baseline: historical open rate by course, section and time in cycle. 3. Add an API endpoint that returns a probability for a section. 4. Build the simple web page on top of it. 5. Deploy and record p95 latency. | A section lookup returns a probability. |
| After spring add/drop | 1. Retrain on a full registration and add/drop cycle. 2. Compare predictions to observed opens and set a calibration target. 3. Try logistic regression only if the baseline is clearly weak. | Calibration is measured on a full cycle and the target is set. |

## Open questions and risks

- [ ] What polling interval does the Rutgers data source tolerate? Unknown. Start slow, watch for errors or blocks, and back off if any appear.
- [ ] Does historical open/close data exist anywhere? Unknown. If not, the model starts from the builder's own logs only, so early forecasts stay rough.
- [ ] Could the data source change format or go offline? Possible. Keep the parser isolated, alert on parse failures, and keep raw snapshots for replay.
- [ ] Will two weeks of data support a useful baseline? Unlikely to be precise. Label the output as a rough baseline and show the history beside it.
- [ ] Will students find and use it during spring registration? Unknown. Share it in Rutgers student communities before registration opens and track lookups.
- [ ] Does a free-tier host keep the poller running unattended? Uncertain. Fall back to a small paid VM if the host sleeps or throttles.
