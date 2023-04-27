The [SLURM REST API](https://slurm.schedmd.com/rest_api.html)
provides a simple mechanism for working with job scripts in
the batch queue.

This project aims to wrap this REST API into the following
file structure,

/{user}/
  pending/{job_id}/
    meta
    script

  running/{job_id}/
    meta
    script
    info

  completed/YYYY/MM/dd/{job_id}
    meta
    script
    info

  new/

  allocations/{proj}/
    info

  status/
    info

Copying a script file into /new will add it to the queue.
Listing entries inside queue, running, or recent will
show information about jobs.  Writing to meta will update
the job queue information.  Job-scripts are read-only, since
slurm does not support updating them.

The allocations directory displays information about
projects and their system usage info.

The status directory displays overall queue and partition
status.

At present, it just wraps a few API calls into golang
functions.
